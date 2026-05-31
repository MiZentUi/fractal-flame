package com.fractalflame.gateway.exception;

public class SseException extends RuntimeException {

    public SseException(String message) {
        super(message);
    }

    public SseException(Throwable cause) {
        super(cause);
    }
}
