package com.fractalflame.gateway.controller;

import org.springframework.core.io.Resource;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.fractalflame.gateway.api.ImagesApi;
import com.fractalflame.gateway.service.ImagesService;

import lombok.RequiredArgsConstructor;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class ImagesController implements ImagesApi {
    private final ImagesService service;

    @Override
    public ResponseEntity<Resource> getImage(String name) {
        return ResponseEntity.ok(service.getByName(name));
    }
}
