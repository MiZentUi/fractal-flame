package com.fractalflame.generator.handler;

import org.jspecify.annotations.Nullable;
import org.springframework.grpc.server.exception.GrpcExceptionHandler;
import org.springframework.stereotype.Component;

import com.fractalflame.generator.exception.BlobNotFoundException;
import com.fractalflame.generator.exception.EntityNotFoundException;
import com.fractalflame.generator.exception.FractalParametersException;
import com.google.rpc.Code;
import com.google.rpc.Status;

import io.grpc.StatusException;
import io.grpc.protobuf.StatusProto;
import jakarta.validation.ValidationException;

@Component
public class GlobalGrpcExceptionHandler implements GrpcExceptionHandler {

    @Override
    public @Nullable StatusException handleException(Throwable exception) {
        var code = switch (exception) {
            case EntityNotFoundException e -> Code.NOT_FOUND;
            case BlobNotFoundException e -> Code.NOT_FOUND;
            case ValidationException e -> Code.INVALID_ARGUMENT;
            case FractalParametersException e -> Code.INVALID_ARGUMENT;
            default -> Code.INTERNAL;
        };

        var status = Status.newBuilder()
                .setCode(code.getNumber())
                .setMessage(exception.getMessage())
                .build();

        return StatusProto.toStatusException(status);
    }
}