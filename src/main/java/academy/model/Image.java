package academy.model;

import academy.cli.AppConfig;

public class Image {
    private final int width;
    private final int height;
    private final Pixel[][] pixels;

    public Image(int width, int height) {
        this.width = width;
        this.height = height;
        pixels = new Pixel[height][width];
    }

    public Image(AppConfig.Size size) {
        width = size.width();
        height = size.height();
        pixels = new Pixel[height][width];
    }

    public int getWidth() {
        return width;
    }

    public int getHeight() {
        return height;
    }

    public Pixel[][] getPixels() {
        return pixels;
    }

    public Pixel getPixel(int x, int y) {
        return pixels[y][x];
    }
}
