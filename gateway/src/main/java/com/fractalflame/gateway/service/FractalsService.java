package com.fractalflame.gateway.service;

import java.util.List;

import org.springframework.stereotype.Service;

import com.fractalflame.gateway.mapper.FractalMapper;
import com.fractalflame.gateway.mapper.TaskMapper;
import com.fractalflame.gateway.model.FractalRequest;
import com.fractalflame.gateway.model.FractalResponse;
import com.fractalflame.gateway.model.TaskState;
import com.fractalflame.generator.proto.IdRequest;
import com.fractalflame.generator.proto.FractalsGrpc.FractalsBlockingStub;
import com.google.protobuf.Empty;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class FractalsService {
    private final FractalsBlockingStub stub;
    private final FractalMapper mapper;
    private final TaskMapper taskMapper;

    public List<FractalResponse> getAll(Integer page, Integer count, String sort, String order, Long userId) {
        var response = stub.getAll(mapper.toFractalsRequest(page, count, sort, order, userId));
        return mapper.toFractalResponseList(response.getFractalsList());
    }

    public FractalResponse getById(Long id) {
        return mapper.toFractalResponse(stub.getById(IdRequest.newBuilder().setId(id).build()));
    }

    public TaskState generation(FractalRequest request) {
        return taskMapper.toTaskState(stub.generation(mapper.fromFractalRequest(request)));
    }

    public List<String> getFunctions() {
        var response = stub.getFunctions(Empty.newBuilder().build());
        return response.getNamesList();
    }
}
