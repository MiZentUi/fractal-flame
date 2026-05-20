package com.fractalflame.generator.service;

import org.springframework.stereotype.Service;

import com.fractalflame.generator.proto.ImageRequest;
import com.fractalflame.generator.proto.ImageResponse;
import com.fractalflame.generator.proto.ImagesGrpc.ImagesImplBase;
import com.google.protobuf.ByteString;

import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class ImagesService extends ImagesImplBase {
    private final StorageService storageService;

    @Override
    public void getImage(ImageRequest request, StreamObserver<ImageResponse> responseObserver) {
        var response = ImageResponse.newBuilder()
                .setImage(ByteString.copyFrom(storageService.getImageBytes(request.getName())))
                .build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
