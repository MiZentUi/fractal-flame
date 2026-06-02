package com.fractalflame.generator.service;

import org.apache.commons.collections4.IteratorUtils;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Sort;
import org.springframework.grpc.server.service.GrpcService;

import com.fractalflame.generator.exception.EntityNotFoundException;
import com.fractalflame.generator.exception.FractalParametersException;
import com.fractalflame.generator.mapper.FractalMapper;
import com.fractalflame.generator.properties.GeneratorProperties;
import com.fractalflame.generator.proto.FractalRequest;
import com.fractalflame.generator.proto.FractalResponse;
import com.fractalflame.generator.proto.FractalsGrpc.FractalsImplBase;
import com.fractalflame.generator.repository.AffineParamsRepository;
import com.fractalflame.generator.repository.FractalsRepository;
import com.fractalflame.generator.repository.FunctionsRepository;
import com.fractalflame.generator.utils.FunctionBuilder;
import com.fractalflame.generator.utils.GenerationTask;
import com.google.protobuf.ByteString;
import com.google.protobuf.Empty;
import com.fractalflame.generator.proto.FractalsRequest;
import com.fractalflame.generator.proto.FractalsResponse;
import com.fractalflame.generator.proto.FunctionsResponse;
import com.fractalflame.generator.proto.IdRequest;
import com.fractalflame.generator.proto.ImageRequest;
import com.fractalflame.generator.proto.ImageResponse;
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
    private final GeneratationService generatationService;
    private final JwtService jwtService;
    private final StorageService storageService;
    private final FractalMapper fractalMapper;
    private final GeneratorProperties generatorProperties;

    @Override
    @Transactional
    public void getAll(FractalsRequest request, StreamObserver<FractalsResponse> responseObserver) {
        var sort = request.hasSort() ? Sort.by(request.getSort()) : Sort.by("id");
        if (request.hasOrder() && request.getOrder() == Order.DESC) {
            sort = sort.descending();
        }

        if (request.hasCount()) {
            var page = request.hasPage() ? request.getPage() : 1;
            var pageRequest = PageRequest.of(page - 1, request.getCount(), sort);
            var fractals = request.hasUserId()
                    ? fractalsRepository.findAllByUserId(request.getUserId(), pageRequest)
                    : fractalsRepository.findAll(pageRequest);
            responseObserver.onNext(FractalsResponse.newBuilder()
                    .addAllFractals(fractalMapper.toFractalResponseList(IteratorUtils.toList(fractals.iterator())))
                    .setPage(page)
                    .setPageCount(fractals.getTotalPages())
                    .build());
        } else {
            var fractals = request.hasUserId()
                    ? fractalsRepository.findAllByUserId(request.getUserId(), sort)
                    : fractalsRepository.findAll(sort);
            responseObserver.onNext(FractalsResponse.newBuilder()
                    .addAllFractals(fractalMapper.toFractalResponseList(IteratorUtils.toList(fractals.iterator())))
                    .setPage(1)
                    .setPageCount(1)
                    .build());
        }
        responseObserver.onCompleted();
    }

    @Override
    @Transactional
    public void getById(IdRequest request, StreamObserver<FractalResponse> responseObserver) {
        responseObserver.onNext(fractalMapper.toFractalResponse(fractalsRepository.findById(request.getId())
                .orElseThrow(() -> new EntityNotFoundException(
                        String.format("Fractal with id = %s not found!", request.getId())))));
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
        var fractal = fractalMapper.fromFractalRequest(request);

        if (fractal.getWidth() > generatorProperties.getMaxWidth()) {
            throw new FractalParametersException(
                    String.format("Fractal width shouldn't be greater than %s!",
                            generatorProperties.getMaxWidth()));
        }

        if (fractal.getHeight() > generatorProperties.getMaxHeight()) {
            throw new FractalParametersException(
                    String.format("Fractal height shouldn't be greater than %s!",
                            generatorProperties.getMaxHeight()));
        }

        if (fractal.getIterationCount() > generatorProperties.getMaxIterations()) {
            throw new FractalParametersException(
                    String.format("Iterations shouldn't be greater than %s!",
                            generatorProperties.getMaxIterations()));
        }

        if (fractal.getSymmetryLevel() > generatorProperties.getMaxSymmetryLevel()) {
            throw new FractalParametersException(String.format("Symmetry level shouldn't be greater than %s!",
                    generatorProperties.getMaxSymmetryLevel()));
        }

        if (fractal.getFunctions().isEmpty()) {
            throw new FractalParametersException("Required at most one function!");
        }

        var functionNames = FunctionBuilder.getFunctionsNames();
        fractal.getFunctions().forEach(f -> {
            if (!functionNames.contains(f.getName())) {
                throw new FractalParametersException("Function with \"" + f.getName() + "\" name not found!");
            }
            if (f.getWeight() <= 0) {
                throw new FractalParametersException("Function weight should be positive!");
            }
        });

        if (fractal.getAffineParams().isEmpty()) {
            throw new FractalParametersException("Required at most one affine params!");
        }

        fractalsRepository.save(fractal);

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

    @Override
    public void subscribeToTask(IdRequest request, StreamObserver<TaskState> responseObserver) {
        generatationService.subscribeToTask(request.getId(), responseObserver);
    }

    @Override
    public void getImage(ImageRequest request, StreamObserver<ImageResponse> responseObserver) {
        var response = ImageResponse.newBuilder()
                .setImage(ByteString.copyFrom(storageService.getImageBytes(request.getName())))
                .build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
