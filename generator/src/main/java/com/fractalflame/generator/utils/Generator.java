package com.fractalflame.generator.utils;

import com.fractalflame.generator.entity.Fractal;
import com.fractalflame.generator.exception.FractalParametersException;
import com.fractalflame.generator.mapper.AffineMapper;
import com.fractalflame.generator.mapper.FunctionMapper;
import com.fractalflame.generator.model.Image;
import com.fractalflame.generator.properties.GeneratorProperties;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

@Slf4j
@RequiredArgsConstructor
public class Generator {
    private final GeneratorProperties properties;
    private final FunctionMapper functionMapper;
    private final AffineMapper affineMapper;
    private Image image;

    public Image getImage() {
        return image;
    }

    public void render(Fractal fractal) {
        log.info("Start render...");

        if (fractal.getSymmetryLevel() < 1) {
            throw new FractalParametersException("WRONG_SYMMETRY_LEVEL", "Symmetry level should be greater than 0!");
        }

        image = new Image(fractal.getWidth(), fractal.getHeight());

        try (var executor = Executors.newFixedThreadPool(properties.getThreads())) {
            for (int i = 0; i < properties.getThreads(); i++) {
                executor.execute(new Render(image,
                        fractal.getIterationCount() / properties.getThreads(),
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
            log.info("Render complete!");
        }

        correction(fractal.getGamma());
    }

    private void correction(Double gamma) {
        log.info("Start correction...");

        var pixels = image.getPixels();

        double max = 0;

        for (var row : pixels) {
            for (var pixel : row) {
                if (pixel != null) {
                    pixel.setNormal(Math.min(Math.log10(pixel.getCounter().sum()), max));
                }
            }
        }

        if (max > 0) {
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

        log.info("Correction complete!");
    }
}
