package academy.utils;

import academy.model.Image;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;
import java.nio.file.Path;

public class ImageWriter {
    private static final Logger LOGGER = LoggerFactory.getLogger(ImageWriter.class);

    private Path outputPath;

    public ImageWriter(Path outputPath) {
        this.outputPath = outputPath;
    }

    public void setOutputPath(Path outputPath) {
        this.outputPath = outputPath;
    }

    public void write(Image image) {
        LOGGER.info("Writing file...");
        System.out.println("Writing file...");
        int width = image.getWidth();
        int height = image.getHeight();
        var pixels = image.getPixels();
        var bufferedImage = new BufferedImage(width, height, BufferedImage.TYPE_INT_RGB);
        for (int i = 0; i < height; i++) {
            for (int j = 0; j < width; j++) {
                if (pixels[i][j] != null){
                    bufferedImage.setRGB(j, i, pixels[i][j].toRGB());
                }
            }
        }
        var file = new File(outputPath.toUri());
        try {
            ImageIO.write(bufferedImage, "png", file);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
