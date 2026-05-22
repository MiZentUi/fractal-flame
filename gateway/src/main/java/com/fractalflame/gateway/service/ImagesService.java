package com.fractalflame.gateway.service;

import org.springframework.core.io.ByteArrayResource;
import org.springframework.core.io.Resource;
import org.springframework.stereotype.Service;

import com.fractalflame.generator.proto.ImageRequest;
import com.fractalflame.generator.proto.ImagesGrpc.ImagesBlockingStub;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class ImagesService {
    private final ImagesBlockingStub stub;

    public Resource getByName(String name) {
        return new ByteArrayResource(
                stub.getImage(ImageRequest.newBuilder().setName(name).build()).getImage().toByteArray());
    }
}
