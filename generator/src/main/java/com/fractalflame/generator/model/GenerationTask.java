package com.fractalflame.generator.model;

import com.fractalflame.generator.entity.Fractal;
import com.fractalflame.generator.proto.TaskState;
import com.fractalflame.generator.utils.ImageWriter;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@Getter
public class GenerationTask {
    private Fractal fractal;
    private Image image;

    @Getter(AccessLevel.NONE)
    private Double progress = 0.0;

    public GenerationTask(Fractal fractal) {
        this.fractal = fractal;
        image = new Image(fractal.getWidth(), fractal.getHeight());
    }

    public synchronized void addProgress(Double value) {
        log.atInfo().addKeyValue("progress", progress).log("add progress");
        progress += value;
        if (Math.abs(progress - 1) >= 0.01) {
            progress = 1.0;
        }
        if ((progress - value) % 10 != progress % 10 || progress == 1) {
            sendUpdate();
        }
    }

    public synchronized void done() {
        progress = 1.0;
        sendUpdate();
    }

    public synchronized void sendUpdate() {
        log.info("send task update");
    }

    public synchronized TaskState toTaskState() {
        synchronized (image) {
            return TaskState.newBuilder()
                    .setFractalId(fractal.getId())
                    .setProgress(progress)
                    .setPreview(ImageWriter.toBase64(image))
                    .build();
        }
    }
}
