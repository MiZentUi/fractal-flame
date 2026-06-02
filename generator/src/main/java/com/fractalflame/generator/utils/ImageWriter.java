package com.fractalflame.generator.utils;

import com.fractalflame.generator.exception.ImageException;
import com.fractalflame.generator.model.Image;
import java.awt.image.BufferedImage;
import java.awt.image.DataBuffer;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.Base64;

import javax.imageio.ImageIO;
import javax.imageio.ImageWriteParam;

public class ImageWriter {
    private ImageWriter() {
    }

    private static final int COMPRESSION_BOUND = 4_194_304;
    private static final float COMPRESSION_QUALITY = 0.75f;

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
            try (var ios = ImageIO.createImageOutputStream(byteStream)) {
                var writers = ImageIO.getImageWritersByFormatName("png");
                if (!writers.hasNext()) {
                    throw new IllegalStateException("No PNG writers found!");
                }
                var writer = writers.next();
                writer.setOutput(ios);

                var param = writer.getDefaultWriteParam();
                if (param.canWriteCompressed()) {
                    param.setCompressionMode(ImageWriteParam.MODE_EXPLICIT);
                    var dataBuffer = bufferedImage.getData().getDataBuffer();
                    var size = dataBuffer.getSize() * DataBuffer.getDataTypeSize(dataBuffer.getDataType()) / 8;
                    if (size > COMPRESSION_BOUND) {
                        param.setCompressionQuality(COMPRESSION_QUALITY);
                    }
                    param.setCompressionQuality(1);
                }
                writer.write(bufferedImage);
            }
            return byteStream.toByteArray();
        } catch (IOException e) {
            throw new ImageException(e);
        }
    }

    public static String toBase64(Image image) {
        return Base64.getEncoder().encodeToString(toByteArray(image));
    }
}
