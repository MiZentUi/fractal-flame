package com.fractalflame.generator.utils;

import com.fractalflame.generator.exception.FractalParametersException;
import com.fractalflame.generator.mapper.AffineMapper;
import com.fractalflame.generator.mapper.FunctionMapper;
import com.fractalflame.generator.model.GenerationTask;
import com.fractalflame.generator.model.Image;
import com.fractalflame.generator.properties.GeneratorProperties;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

import org.springframework.stereotype.Component;

@Component
@Slf4j
@RequiredArgsConstructor
public class Generator {
    private final GeneratorProperties properties;
    private final FunctionMapper functionMapper;
    private final AffineMapper affineMapper;

    public void process(GenerationTask task) {
        var fractal = task.getFractal();

        log.atInfo().addKeyValue("fractal_id", fractal.getId()).log("Start render...");

        try (var executor = Executors.newFixedThreadPool(properties.getThreads())) {
            for (int i = 0; i < properties.getThreads(); i++) {
                executor.execute(new Render(task,
                        fractal.getIterationCount() / properties.getThreads(),
                        fractal.getIterationCount(),
                        properties.getThreads(),
                        fractal.getSymmetryLevel(),
                        functionMapper.toFunctionModelList(fractal.getFunctions()),
                        affineMapper.toAffineTransformationList(fractal.getAffineParams())));
            }
            executor.shutdown();
            try {
                if (!executor.awaitTermination(30, TimeUnit.MINUTES)) {
                    executor.shutdownNow();
                }
            } catch (InterruptedException e) {
                executor.shutdownNow();
                Thread.currentThread().interrupt();
            }
            log.atInfo().addKeyValue("fractal_id", fractal.getId()).log("Render complete!");
        }

        task.sendUpdate();
        correction(task.getImage(), fractal.getGamma());
        task.done();
    }

    private void correction(Image image, Double gamma) {
        log.info("Start correction...");

        var pixels = image.getPixels();

        double max = 0;

        for (var row : pixels) {
            for (var pixel : row) {
                if (pixel != null) {
                    pixel.setNormal(Math.log10(pixel.getCounter().sum()));
                    max = Math.max(pixel.getNormal(), max);
                }
            }
        }

        if (max <= 0) {
            return;
        }

        for (var row : pixels) {
            for (var pixel : row) {
                if (pixel != null) {
                    var normal = pixel.getNormal();
                    normal /= max;
                    pixel.setNormal(normal);
                    pixel.multiply(Math.pow(normal, 1.0 / gamma));
                }
            }
        }
    }
}
