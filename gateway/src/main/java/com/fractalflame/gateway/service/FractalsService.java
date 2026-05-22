package com.fractalflame.gateway.service;

import java.util.List;

import org.springframework.stereotype.Service;
import org.springframework.web.servlet.mvc.method.annotation.SseEmitter;

import com.fractalflame.gateway.mapper.FractalMapper;
import com.fractalflame.gateway.mapper.TaskMapper;
import com.fractalflame.gateway.model.FractalRequest;
import com.fractalflame.gateway.model.FractalResponse;
import com.fractalflame.generator.proto.IdRequest;
import com.fractalflame.generator.proto.TaskState;
import com.fractalflame.generator.proto.FractalsGrpc.FractalsBlockingStub;
import com.fractalflame.generator.proto.FractalsGrpc.FractalsStub;
import com.google.protobuf.Empty;

import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class FractalsService {
    private final FractalsBlockingStub blockingStub;
    private final FractalsStub stub;
    private final FractalMapper mapper;
    private final TaskMapper taskMapper;

    public List<FractalResponse> getAll(Integer page, Integer count, String sort, String order, Long userId) {
        var response = blockingStub.getAll(mapper.toFractalsRequest(page, count, sort, order, userId));
        return mapper.toFractalResponseList(response.getFractalsList());
    }

    public FractalResponse getById(Long id) {
        return mapper.toFractalResponse(blockingStub.getById(IdRequest.newBuilder().setId(id).build()));
    }

    public com.fractalflame.gateway.model.TaskState generation(FractalRequest request) {
        return taskMapper.toTaskState(blockingStub.generation(mapper.fromFractalRequest(request)));
    }

    public void subscribeToTask(Long id, SseEmitter emitter) {
        stub.subscribeToTask(IdRequest.newBuilder().setId(id).build(), new StreamObserver<TaskState>() {

            @Override
            public void onNext(TaskState value) {
                try {
                    emitter.send(SseEmitter.event()
                            .data(taskMapper.toTaskState(value)));

                    if (Math.abs(value.getProgress() - 1) < 0.005) {
                        emitter.complete();
                    }
                } catch (Exception e) {
                    emitter.completeWithError(e);
                }
            }

            @Override
            public void onError(Throwable t) {
                emitter.completeWithError(t);
            }

            @Override
            public void onCompleted() {
                emitter.complete();
            }

        });
    }

    public List<String> getFunctions() {
        var response = blockingStub.getFunctions(Empty.newBuilder().build());
        return response.getNamesList();
    }
}
