package com.fractalflame.generator.exception;

public class FunctionException extends RuntimeException {

    public FunctionException(String message) {
        super(message);
    }

    public FunctionException(Throwable cause) {
        super(cause);
    }
}
