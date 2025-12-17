package academy.utils;

import academy.cli.AppConfig;
import academy.model.AffineTransformation;
import academy.model.Image;
import academy.model.Pixel;
import academy.model.Point;
import academy.model.functions.Function;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import java.util.Comparator;
import java.util.List;
import java.util.Random;

public class Render implements Runnable {
    private final static Logger LOGGER = LoggerFactory.getLogger(Render.class);

    private final Random random;
    private final int iterationsCount;
    private final List<Function> functions;
    private final List<AffineTransformation> affineTransformations;
    private final Image image;
    private final int symmetryLevel;

    public Render(Image image, AppConfig config, int iterationsCount) {
        this.image = image;
        this.iterationsCount = iterationsCount;
        functions = config.getFunctions();
        AffineTransformation.setSeed(config.getSeed());
        affineTransformations = config.getAffineParams().stream().map(AffineTransformation::new).toList();
        random = new Random(config.getSeed() + Thread.currentThread().threadId());
        symmetryLevel = config.getSymmetryLevel();
    }

    @Override
    public void run() {
        synchronized (LOGGER) {
            LOGGER.info("[Thread: {}] Start generation...", this.hashCode());
        }

        var width = image.getWidth();
        var height = image.getHeight();
        var pixels = image.getPixels();

        double x_max = width > height ? (double) width / height : 1;
        double y_max = width < height ? (double) height / width : 1;
        double x_min = -x_max, y_min = -y_max;
        double theta = 2 * Math.PI / symmetryLevel;

        var point = new Point(random.nextDouble(x_min, x_max), random.nextDouble(y_min, y_max));
        var color = Pixel.genColor(random);

        for (int i = -20; i < iterationsCount; i++) {
            synchronized (System.out) {
                System.out.print('\r' + String.format("[Thread: %d] Generating: %.2f%%", this.hashCode(), (double) i / (iterationsCount + 20) * 100));
            }

            var currentAffine = getRandomAffine();
            point = getRandomFunction().transform(currentAffine.transform(point));

            for (int j = 0; j < symmetryLevel; j++) {
                double angle = j * theta;

                double x_rot = point.x * Math.cos(angle) - point.y * Math.sin(angle);
                double y_rot = point.x * Math.sin(angle) + point.y * Math.cos(angle);

                point = new Point(x_rot, y_rot);

                int x = width - (int) ((x_max - point.x) / (x_max - x_min) * width);
                int y = height - (int) ((y_max - point.y) / (y_max - y_min) * height);

                color.add(currentAffine.getColor()).multiply(0.5);

                if (i >= 0 && 0 < x && x < width && 0 < y && y < height) {
                    var pixel = new Pixel(color);
                    if (pixels[y][x] != null) {
                        pixel.counter = pixels[y][x].counter;
                    }
                    pixels[y][x] = pixel;
                    pixels[y][x].counter++;
                }
            }
        }

        synchronized (LOGGER) {
            LOGGER.info("[Thread: {}] Generation complete!", this.hashCode());
        }
    }

    private AffineTransformation getRandomAffine() {
        return affineTransformations.get(random.nextInt(affineTransformations.size()));
    }

    private Function getRandomFunction() {
        double sum = functions.stream().mapToDouble(Function::getWidth).sum();
        double n = 0;
        double r_num = random.nextDouble(sum);
        for (var function : functions) {
            n += function.getWidth();
            if (n >= r_num) {
                return function;
            }
        }
        return functions.stream().max(Comparator.comparing(Function::getWidth)).orElseGet(functions::getFirst);
    }
}
