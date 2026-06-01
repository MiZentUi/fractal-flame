package com.fractalflame.gateway.service;

import java.io.IOException;
import java.util.List;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.io.ByteArrayResource;
import org.springframework.core.io.Resource;
import org.springframework.stereotype.Service;
import org.springframework.web.servlet.mvc.method.annotation.SseEmitter;

import com.fractalflame.gateway.exception.SseException;
import com.fractalflame.gateway.mapper.FractalMapper;
import com.fractalflame.gateway.mapper.TaskMapper;
import com.fractalflame.gateway.model.FractalRequest;
import com.fractalflame.gateway.model.FractalResponse;
import com.fractalflame.gateway.model.FractalsResponse;
import com.fractalflame.generator.proto.IdRequest;
import com.fractalflame.generator.proto.ImageRequest;
import com.fractalflame.generator.proto.TaskState;
import com.fractalflame.generator.proto.FractalsGrpc.FractalsBlockingStub;
import com.fractalflame.generator.proto.FractalsGrpc.FractalsStub;
import com.google.protobuf.Empty;

import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@Service
@RequiredArgsConstructor
@Slf4j
public class FractalsService {
    private final FractalsBlockingStub blockingStub;
    private final FractalsStub stub;
    private final FractalMapper mapper;
    private final TaskMapper taskMapper;

    @Value("${app.sse-timeout}")
    private Long sseTimeout;

    public FractalsResponse getAll(Integer page, Integer count, String sort, String order, Long userId) {
        var response = blockingStub.getAll(mapper.toFractalsRequest(page, count, sort, order, userId));
        return FractalsResponse.builder()
                .items(mapper.toFractalResponseList(response.getFractalsList()))
                .page(response.getPage())
                .pageCount(response.getPageCount())
                .build();
    }

    public FractalResponse getById(Long id) {
        return mapper.toFractalResponse(blockingStub.getById(IdRequest.newBuilder().setId(id).build()));
    }

    public com.fractalflame.gateway.model.TaskState generation(FractalRequest request) {
        return taskMapper.toTaskState(blockingStub.generation(mapper.fromFractalRequest(request)));
    }

    public SseEmitter subscribeToTask(Long id) {
        var emitter = new SseEmitter(sseTimeout);

        stub.subscribeToTask(IdRequest.newBuilder().setId(id).build(), new StreamObserver<TaskState>() {

            @Override
            public void onNext(TaskState value) {
                try {
                    log.atInfo().addKeyValue("fractal_id", value.getFractalId()).log("send update over sse");
                    emitter.send(SseEmitter.event()
                            .name("update")
                            .data(taskMapper.toTaskState(value)));

                    if (Math.abs(value.getProgress() - 1) < 0.005) {
                        log.atInfo().addKeyValue("fractal_id", value.getFractalId()).log("emitter complete");

                        emitter.send(SseEmitter.event()
                                .name("complete")
                                .data("generation complete"));

                        emitter.complete();
                    }
                } catch (Exception e) {
                    sendError(e);
                    emitter.completeWithError(new SseException(e));
                }
            }

            @Override
            public void onError(Throwable t) {
                sendError(t);
                emitter.completeWithError(new SseException(t));
            }

            @Override
            public void onCompleted() {
                try {
                    emitter.send(SseEmitter.event()
                            .name("complete"));
                } catch (Exception e) {
                    sendError(e);
                    emitter.completeWithError(new SseException(e));
                }
                emitter.complete();
            }

            private void sendError(Throwable throwable) {
                try {
                    emitter.send(SseEmitter.event()
                            .name("error")
                            .data(throwable.getMessage()));
                } catch (IOException ex) {
                    // dropped connection
                }
            }

        });

        return emitter;
    }

    public List<String> getFunctions() {
        var response = blockingStub.getFunctions(Empty.newBuilder().build());
        return response.getNamesList();
    }

    public Resource getImageByName(String name) {
        return new ByteArrayResource(blockingStub.getImage(ImageRequest.newBuilder()
                .setName(name)
                .build()).getImage().toByteArray());
    }
}
