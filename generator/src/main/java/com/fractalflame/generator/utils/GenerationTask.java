package com.fractalflame.generator.utils;

import java.util.concurrent.atomic.AtomicLong;

import com.fractalflame.generator.entity.Fractal;
import com.fractalflame.generator.model.Image;
import com.fractalflame.generator.proto.TaskState;

import io.grpc.stub.StreamObserver;
import lombok.AccessLevel;
import lombok.Getter;
import lombok.Setter;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@Getter
public class GenerationTask {
    private Fractal fractal;
    private Image image;

    @Setter
    private StreamObserver<TaskState> responseObserver;

    @Getter(AccessLevel.NONE)
    private final AtomicLong iterations = new AtomicLong();

    public GenerationTask(Fractal fractal) {
        this.fractal = fractal;
        image = new Image(fractal.getWidth(), fractal.getHeight());
    }

    public void addIterations(Long iterations) {
        this.iterations.addAndGet(iterations);
    }

    public synchronized TaskState toTaskState() {
        var progress = (double) iterations.get() / (fractal.getIterationCount() - 10);
        if (progress > 1) {
            progress = 1.0;
        }
        synchronized (image) {
            return TaskState.newBuilder()
                    .setFractalId(fractal.getId())
                    .setProgress(progress)
                    .setPreview(ImageWriter.toBase64(image))
                    .build();
        }
    }
}
