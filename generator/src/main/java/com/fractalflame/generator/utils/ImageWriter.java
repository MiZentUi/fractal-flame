package com.fractalflame.generator.utils;

import com.fractalflame.generator.model.Image;
import java.awt.image.BufferedImage;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.Base64;

import javax.imageio.ImageIO;

public class ImageWriter {
    private ImageWriter() {
    }

    public static byte[] toByteArray(Image image) {
        int width = image.getWidth();
        int height = image.getHeight();
        var pixels = image.getPixels();
        var bufferedImage = new BufferedImage(width, height, BufferedImage.TYPE_INT_RGB);
        for (int i = 0; i < height; i++) {
            for (int j = 0; j < width; j++) {
                if (pixels[i][j] != null) {
                    bufferedImage.setRGB(j, i, pixels[i][j].toRGB());
                }
            }
        }
        try {
            var byteStream = new ByteArrayOutputStream();
            ImageIO.write(bufferedImage, "png", byteStream);
            return byteStream.toByteArray();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public static String toBase64(Image image) {
        return Base64.getEncoder().encodeToString(toByteArray(image));
    }
}
