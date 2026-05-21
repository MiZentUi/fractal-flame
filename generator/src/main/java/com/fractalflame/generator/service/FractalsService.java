package com.fractalflame.generator.service;

import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Sort;
import org.springframework.grpc.server.service.GrpcService;
import com.fractalflame.generator.exception.EntityNotFoundException;
import com.fractalflame.generator.mapper.FractalMapper;
import com.fractalflame.generator.model.GenerationTask;
import com.fractalflame.generator.proto.FractalRequest;
import com.fractalflame.generator.proto.FractalResponse;
import com.fractalflame.generator.proto.FractalsGrpc.FractalsImplBase;
import com.fractalflame.generator.repository.AffineParamsRepository;
import com.fractalflame.generator.repository.FractalsRepository;
import com.fractalflame.generator.repository.FunctionsRepository;
import com.fractalflame.generator.utils.FunctionBuilder;
import com.google.protobuf.Empty;
import com.fractalflame.generator.proto.FractalsRequest;
import com.fractalflame.generator.proto.FractalsResponse;
import com.fractalflame.generator.proto.FunctionsResponse;
import com.fractalflame.generator.proto.IdRequest;
import com.fractalflame.generator.proto.Order;
import com.fractalflame.generator.proto.TaskState;

import io.grpc.stub.StreamObserver;
import jakarta.transaction.Transactional;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@GrpcService
@Slf4j
@RequiredArgsConstructor
public class FractalsService extends FractalsImplBase {
    private final FractalsRepository fractalsRepository;
    private final FunctionsRepository functionsRepository;
    private final AffineParamsRepository affineParamsRepository;
    private final FractalMapper fractalMapper;
    private final JwtService jwtService;
    private final GeneratationService generatationService;

    @Override
    @Transactional
    public void getAll(FractalsRequest request, StreamObserver<FractalsResponse> responseObserver) {
        try {
            var sort = request.hasSort() ? Sort.by(request.getSort()) : Sort.by("id");
            if (request.hasOrder() && request.getOrder() == Order.DESC) {
                sort = sort.descending();
            }
            var page = request.hasPage() ? request.getPage() : 1;
            var count = request.hasCount() ? request.getCount() : 10;
            var pageRequest = PageRequest.of(page - 1, count, sort);
            var fractals = request.hasUserId() ? fractalsRepository.findAllByUserId(request.getUserId(), pageRequest)
                    : fractalsRepository.findAll(pageRequest);

            responseObserver.onNext(FractalsResponse.newBuilder()
                    .addAllFractals(fractalMapper.toFractalResponseList(fractals.toList()))
                    .build());
        } catch (Exception e) {
            responseObserver.onError(e);
            return;
        }
        responseObserver.onCompleted();
    }

    @Override
    @Transactional
    public void getById(IdRequest request, StreamObserver<FractalResponse> responseObserver) {
        try {
            responseObserver.onNext(fractalMapper.toFractalResponse(fractalsRepository.findById(request.getId())
                    .orElseThrow(() -> new EntityNotFoundException(
                            String.format("Fractal with id = %s not found!", request.getId())))));
        } catch (Exception e) {
            responseObserver.onError(e);
            return;
        }
        responseObserver.onCompleted();
    }

    @Override
    public void getFunctions(Empty request, StreamObserver<FunctionsResponse> responseObserver) {
        responseObserver.onNext(FunctionsResponse.newBuilder()
                .addAllNames(FunctionBuilder.getFunctionsNames())
                .build());
        responseObserver.onCompleted();
    }

    @Override
    @Transactional
    public void generation(FractalRequest request, StreamObserver<TaskState> responseObserver) {
        var fractal = fractalsRepository.save(fractalMapper.fromFractalRequest(request));

        var user = jwtService.getCurrentUser();
        if (user != null) {
            fractal.setUserId(user.getId());
        }

        fractal.getFunctions().forEach(f -> {
            f.setFractal(fractal);
            functionsRepository.save(f);
        });

        fractal.getAffineParams().forEach(p -> {
            p.setFractal(fractal);
            affineParamsRepository.save(p);
        });

        var task = new GenerationTask(fractal);

        generatationService.addTask(task);

        responseObserver.onNext(task.toTaskState());
        responseObserver.onCompleted();
    }
}
