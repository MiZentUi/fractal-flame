package academy.utils;

import academy.cli.AppConfig;
import academy.model.Image;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Generator {
    private static final Logger LOGGER = LoggerFactory.getLogger(Generator.class);

    private Image image;
    private final AppConfig config;

    public Generator(AppConfig config) {
        this.config = config;
    }

    public Image getImage() {
        return image;
    }

    public void render() {
        LOGGER.info("Start render...");

        image = new Image(config.getSize());

        try (var executor = Executors.newFixedThreadPool(config.getThreads())) {
            for (int i = 0; i < config.getThreads(); i++) {
                executor.execute(new Render(image, config, config.getIterationCount() / config.getThreads()));
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
            System.out.println();

            LOGGER.info("Render complete!");
            System.out.println("Render complete!");
        }
    }

    public void correction() {
        LOGGER.info("Start correction...");

        var pixels = image.getPixels();
        var gamma = config.getGamma();

        double max = 0;

        for (var row : pixels) {
            for (var pixel : row) {
                if (pixel != null) {
                    pixel.normal = Math.log10(pixel.counter.sum());
                    if (pixel.normal > max) {
                        max = pixel.normal;
                    }
                }
            }
        }

        if (max > 0) {
            for (var row : pixels) {
                for (var pixel : row) {
                    if (pixel != null) {
                        pixel.normal /= max;
                        pixel.multiply(Math.pow(pixel.normal, 1.0 / gamma));
                    }
                }
            }
        }

        LOGGER.info("Correction complete!");
        System.out.println("Correction complete!");
    }
}
