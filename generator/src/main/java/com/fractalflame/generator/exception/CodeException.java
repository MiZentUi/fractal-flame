package com.fractalflame.generator.exception;

import lombok.Getter;

@Getter
public abstract class CodeException extends RuntimeException {
    private final String code;

    protected CodeException(String code, String message) {
        super(message);
        this.code = code;
    }

    protected CodeException(String code, Throwable cause) {
        super(cause);
        this.code = code;
    }
}
