package com.fractalflame.gateway.controller;

import java.util.List;

import org.jspecify.annotations.Nullable;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.fractalflame.gateway.api.FractalsApi;
import com.fractalflame.gateway.model.FractalRequest;
import com.fractalflame.gateway.model.FractalResponse;
import com.fractalflame.gateway.model.TaskState;
import com.fractalflame.gateway.service.FractalsService;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class FractalsController implements FractalsApi {
    private final FractalsService service;

    @Override
    public ResponseEntity<List<FractalResponse>> getFractals(@Valid @Nullable Integer page,
            @Valid @Nullable Integer count,
            @Valid @Nullable String sort, @Valid @Nullable String order, @Valid @Nullable Long userid) {
        return ResponseEntity.ok(service.getAll(page, count, sort, order, userid));
    }

    @Override
    public ResponseEntity<FractalResponse> getFractal(Long id) {
        return ResponseEntity.ok(service.getById(id));
    }

    @Override
    public ResponseEntity<TaskState> generation(@Valid FractalRequest fractalRequest) {
        return ResponseEntity.ok(service.generation(fractalRequest));
    }

    @Override
    public ResponseEntity<List<String>> getFunctions() {
        return ResponseEntity.ok(service.getFunctions());
    }
}
