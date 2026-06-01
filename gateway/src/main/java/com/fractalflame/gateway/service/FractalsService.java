package com.fractalflame.gateway.service;

import java.util.List;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.io.ByteArrayResource;
import org.springframework.core.io.Resource;
import org.springframework.stereotype.Service;
import org.springframework.web.servlet.mvc.method.annotation.SseEmitter;

import com.fractalflame.gateway.mapper.FractalMapper;
import com.fractalflame.gateway.mapper.TaskMapper;
import com.fractalflame.gateway.model.FractalRequest;
import com.fractalflame.gateway.model.FractalResponse;
import com.fractalflame.gateway.model.FractalsResponse;
import com.fractalflame.gateway.service.observer.TaskEventsObserver;
import com.fractalflame.generator.proto.IdRequest;
import com.fractalflame.generator.proto.ImageRequest;
import com.fractalflame.generator.proto.FractalsGrpc.FractalsBlockingStub;
import com.fractalflame.generator.proto.FractalsGrpc.FractalsStub;
import com.google.protobuf.Empty;

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

        stub.subscribeToTask(IdRequest.newBuilder().setId(id).build(),
                new TaskEventsObserver(emitter, taskMapper));

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
