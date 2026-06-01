package com.fractalflame.gateway.service.observer;

import java.io.IOException;
import java.util.concurrent.atomic.AtomicBoolean;

import org.springframework.web.servlet.mvc.method.annotation.SseEmitter;

import com.fractalflame.gateway.exception.SseException;
import com.fractalflame.gateway.mapper.TaskMapper;
import com.fractalflame.generator.proto.IdRequest;
import com.fractalflame.generator.proto.TaskState;

import io.grpc.Status;
import io.grpc.stub.ClientCallStreamObserver;
import io.grpc.stub.ClientResponseObserver;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@RequiredArgsConstructor
public class TaskEventsObserver implements ClientResponseObserver<IdRequest, TaskState> {
    private final SseEmitter emitter;
    private final TaskMapper taskMapper;

    private ClientCallStreamObserver<IdRequest> requestStream;
    private final AtomicBoolean canceled = new AtomicBoolean(false);

    @Override
    public void beforeStart(ClientCallStreamObserver<IdRequest> requestStream) {
        this.requestStream = requestStream;
    }

    private record MessageEvent(String message) {
    }

    @Override
    public void onNext(TaskState value) {
        try {
            log.atInfo()
                    .addKeyValue("fractal_id", value.getFractalId())
                    .addKeyValue("progress", value.getProgress())
                    .log("send update over sse");
            emitter.send(SseEmitter.event()
                    .name("update")
                    .data(taskMapper.toTaskState(value)));

            if (Math.abs(value.getProgress() - 1) < 0.005) {
                log.atInfo()
                        .addKeyValue("fractal_id", value.getFractalId())
                        .log("emitter complete");

                emitter.send(SseEmitter.event()
                        .name("complete")
                        .data(new MessageEvent("generation complete")));

                cancel("complete", null);

                emitter.complete();
            }
        } catch (Exception e) {
            cancel("error", e);
            sendErrorAndComplete(e);
        }
    }

    @Override
    public void onError(Throwable t) {
        log.atWarn().addKeyValue("message", t.getMessage()).log("subscribe response observer on error");
        if (Status.fromThrowable(t).getCode() != Status.Code.CANCELLED || !canceled.get()) {
            sendErrorAndComplete(t);
        }
    }

    @Override
    public void onCompleted() {
        try {
            emitter.send(SseEmitter.event()
                    .name("complete")
                    .data(new MessageEvent("generation complete")));
        } catch (Exception e) {
            sendErrorAndComplete(e);
        }
        emitter.complete();
    }

    private void cancel(String message, Exception exception) {
        canceled.set(true);
        requestStream.cancel(message, exception);
    }

    private void sendErrorAndComplete(Throwable throwable) {
        try {
            emitter.send(SseEmitter.event()
                    .name("error")
                    .data(throwable.getMessage()));
        } catch (IOException ex) {
            // dropped connection
        }
        emitter.completeWithError(new SseException(throwable));
    }
}
