package academy.model;

import java.util.Random;
import java.util.concurrent.atomic.LongAdder;

public class Pixel {
    public int red;
    public int green;
    public int blue;
    public int alpha;
    public LongAdder counter;
    public double normal;

    public Pixel(Pixel pixel) {
        if (pixel != null) {
            this.red = pixel.red;
            this.green = pixel.green;
            this.blue = pixel.blue;
            this.alpha = pixel.alpha;
        }
        counter = new LongAdder();
        normal = 0;
    }

    public Pixel(int red, int green, int blue) {
        this.red = red;
        this.green = green;
        this.blue = blue;
        alpha = 255;
        counter = new LongAdder();
        normal = 0;
    }

    public Pixel(int red, int green, int blue, int alpha) {
        this.red = red;
        this.green = green;
        this.blue = blue;
        this.alpha = alpha;
    }

    public static Pixel add(Pixel pixel1, Pixel pixel2) {
        if (pixel1 == null) {
            return pixel2;
        }
        var pixel = new Pixel(pixel1);
        pixel.counter = pixel1.counter;
        return pixel.add(pixel2);
    }

    public static Pixel unite(Pixel pixel1, Pixel pixel2) {
        var pixel = add(pixel1, pixel2);
        return pixel1 != null && pixel2 != null ? pixel.multiply(0.5) : pixel;
    }

    public Pixel add(Pixel pixel) {
        if (pixel != null) {
            red += pixel.red;
            blue += pixel.blue;
            green += pixel.green;
            counter.add(pixel.counter.sum());
        }
        return this;
    }

    public Pixel multiply(double num) {
        red = (int) Math.round(red * num);
        green = (int) Math.round(green * num);
        blue = (int) Math.round(blue * num);
        return this;
    }

    public int toRGB() {
        return (alpha << 24) | (red << 16) | (green << 8) | blue;
    }

    public static Pixel genColor(Random random) {
        int low = 0, high = 255;
        int red = random.nextInt(low, high);
        int green = random.nextInt(low, high);
        int blue = random.nextInt(low, high);
        return new Pixel(red, green, blue);
    }
}
