package com.fractalflame.gateway.advice;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.HttpMediaTypeNotAcceptableException;
import org.springframework.web.bind.MissingRequestCookieException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.context.request.async.AsyncRequestNotUsableException;
import org.springframework.web.servlet.resource.NoResourceFoundException;

import com.fractalflame.gateway.exception.SseException;
import com.fractalflame.gateway.model.ApiStatusResponse;

import io.grpc.StatusRuntimeException;
import io.jsonwebtoken.JwtException;
import lombok.extern.slf4j.Slf4j;

@RestControllerAdvice
@Slf4j
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

        @ExceptionHandler(JwtException.class)
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

        @ExceptionHandler({ SseException.class, AsyncRequestNotUsableException.class })
        public void handleSse(Exception exception) {
                log.atInfo().addKeyValue("message", exception.getMessage()).log("sse exception");
        }

        @ExceptionHandler({ NoResourceFoundException.class, MissingRequestCookieException.class })
        public ResponseEntity<ApiStatusResponse> handleNotFound(Exception exception) {
                var status = HttpStatus.NOT_FOUND;
                return new ResponseEntity<>(
                                ApiStatusResponse.builder()
                                                .code(status.value())
                                                .status(status.name())
                                                .message(exception.getMessage())
                                                .build(),
                                status);
        }

        @ExceptionHandler(HttpMediaTypeNotAcceptableException.class)
        public ResponseEntity<ApiStatusResponse> handleNotAcceptable(Exception exception) {
                var status = HttpStatus.NOT_ACCEPTABLE;
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
