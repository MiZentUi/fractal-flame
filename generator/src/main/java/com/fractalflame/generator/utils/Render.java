package com.fractalflame.generator.utils;

import com.fractalflame.generator.model.AffineTransformation;
import com.fractalflame.generator.model.Image;
import com.fractalflame.generator.model.Pixel;
import com.fractalflame.generator.model.Point;
import com.fractalflame.generator.model.functions.FunctionModel;

import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import java.util.Comparator;
import java.util.List;
import java.util.Random;

@Slf4j
@AllArgsConstructor
public class Render implements Runnable {
    private final Random random = new Random(System.currentTimeMillis());
    private final Image image;
    private final int iterationsCount;
    private final int symmetryLevel;
    private final List<FunctionModel> functions;
    private final List<AffineTransformation> affineTransformations;

    @Override
    public void run() {
        synchronized (log) {
            log.info("[Thread: {}] Start generation...", this.hashCode());
        }

        var width = image.getWidth();
        var height = image.getHeight();
        var pixels = image.getPixels();

        double xMax = width > height ? (double) width / height : 1;
        double yMax = width < height ? (double) height / width : 1;
        double xMin = -xMax;
        double yMin = -yMax;
        double theta = 2 * Math.PI / symmetryLevel;

        var point = new Point(random.nextDouble(xMin, xMax), random.nextDouble(yMin, yMax));
        var color = Pixel.genColor(random);

        for (int i = -20; i < iterationsCount; i++) {
            var currentAffine = getRandomAffine();
            point = getRandomFunction().transform(currentAffine.transform(point));

            for (int j = 0; j < symmetryLevel; j++) {
                double angle = j * theta;

                double xRot = point.x() * Math.cos(angle) - point.y() * Math.sin(angle);
                double yRot = point.x() * Math.sin(angle) + point.y() * Math.cos(angle);

                point = new Point(xRot, yRot);

                int x = width - (int) ((xMax - point.x()) / (xMax - xMin) * width);
                int y = height - (int) ((yMax - point.y()) / (yMax - yMin) * height);

                color.add(currentAffine.getColor()).multiply(0.5);

                if (i >= 0 && 0 < x && x < width && 0 < y && y < height) {
                    var pixel = new Pixel(color);
                    synchronized (image) {
                        if (pixels[y][x] != null) {
                            pixel.setCounter(pixels[y][x].getCounter());
                        }
                        pixels[y][x] = pixel;
                    }
                    pixels[y][x].getCounter().increment();
                }
            }
        }

        synchronized (log) {
            log.info("[Thread: {}] Generation complete!", this.hashCode());
        }
    }

    private AffineTransformation getRandomAffine() {
        return affineTransformations.get(random.nextInt(affineTransformations.size()));
    }

    private FunctionModel getRandomFunction() {
        double sum = functions.stream().mapToDouble(FunctionModel::getWidth).sum();
        double n = 0;
        double rNum = random.nextDouble(sum);
        for (var function : functions) {
            n += function.getWidth();
            if (n >= rNum) {
                return function;
            }
        }
        return functions.stream().max(Comparator.comparing(FunctionModel::getWidth)).orElseGet(functions::getFirst);
    }
}
