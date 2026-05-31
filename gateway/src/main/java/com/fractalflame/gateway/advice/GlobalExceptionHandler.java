package com.fractalflame.gateway.advice;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import com.fractalflame.gateway.model.ApiStatusResponse;

import io.grpc.StatusRuntimeException;
import io.jsonwebtoken.ExpiredJwtException;
import io.jsonwebtoken.security.SignatureException;

@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(StatusRuntimeException.class)
    public ResponseEntity<ApiStatusResponse> handleStatusRuntimeException(StatusRuntimeException exception) {
        var status = exception.getStatus();

        var httpStatus = switch (status.getCode()) {
            case NOT_FOUND -> HttpStatus.NOT_FOUND;
            case INVALID_ARGUMENT -> HttpStatus.BAD_REQUEST;
            case UNAUTHENTICATED -> HttpStatus.UNAUTHORIZED;
            case PERMISSION_DENIED -> HttpStatus.FORBIDDEN;
            case DEADLINE_EXCEEDED -> HttpStatus.GATEWAY_TIMEOUT;
            case ALREADY_EXISTS -> HttpStatus.CONFLICT;
            default -> HttpStatus.INTERNAL_SERVER_ERROR;
        };

        return new ResponseEntity<>(
                ApiStatusResponse.builder()
                        .code(httpStatus.value())
                        .status(status.getCode().name())
                        .message(exception.getMessage())
                        .build(),
                httpStatus);
    }

    @ExceptionHandler({ SignatureException.class, ExpiredJwtException.class })
    public ResponseEntity<ApiStatusResponse> handleSigning(Exception exception) {
        var status = HttpStatus.UNAUTHORIZED;
        return new ResponseEntity<>(
                ApiStatusResponse.builder()
                        .code(status.value())
                        .status(status.name())
                        .message(exception.getMessage())
                        .build(),
                status);
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<ApiStatusResponse> handle(Exception exception) {
        var status = HttpStatus.INTERNAL_SERVER_ERROR;
        return new ResponseEntity<>(
                ApiStatusResponse.builder()
                        .code(status.value())
                        .status(status.name())
                        .message(exception.getMessage())
                        .build(),
                status);
    }
}
