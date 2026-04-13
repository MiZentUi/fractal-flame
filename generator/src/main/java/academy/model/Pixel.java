package academy.model;

import java.util.Random;
import java.util.concurrent.atomic.LongAdder;

public class Pixel {
    public LongAdder red = new LongAdder();
    public LongAdder green = new LongAdder();
    public LongAdder blue = new LongAdder();
    public LongAdder alpha = new LongAdder();
    public LongAdder counter = new LongAdder();
    public double normal;

    public Pixel(Pixel pixel) {
        if (pixel != null) {
            this.red.add(pixel.red.sum());
            this.green.add(pixel.green.sum());
            this.blue.add(pixel.blue.sum());
            this.alpha.add(pixel.alpha.sum());
        }
        counter = new LongAdder();
        normal = 0;
    }

    public Pixel(int red, int green, int blue) {
        this.red.add(red);
        this.green.add(green);
        this.blue.add(blue);
        alpha.add(255);
        counter = new LongAdder();
        normal = 0;
    }

    public Pixel(int red, int green, int blue, int alpha) {
        this.red.add(red);
        this.green.add(green);
        this.blue.add(blue);
        this.alpha.add(alpha);
    }

    public static Pixel add(Pixel pixel1, Pixel pixel2) {
        if (pixel1 == null) {
            return pixel2;
        }
        var pixel = new Pixel(pixel1);
        pixel.counter.add(pixel1.counter.sum());
        return pixel.add(pixel2);
    }

    public Pixel add(Pixel pixel) {
        if (pixel != null) {
            this.red.add(pixel.red.sum());
            this.green.add(pixel.green.sum());
            this.blue.add(pixel.blue.sum());
            counter.add(pixel.counter.sum());
        }
        return this;
    }

    public Pixel multiply(double num) {
        red.add(Math.round(red.sum() * (num - 1)));
        green.add(Math.round(green.sum() * (num - 1)));
        blue.add(Math.round(blue.sum() * (num - 1)));
        return this;
    }

    public int toRGB() {
        int alpha = Math.toIntExact(this.alpha.sum());
        int red = Math.toIntExact(this.red.sum());
        int green = Math.toIntExact(this.green.sum());
        int blue = Math.toIntExact(this.blue.sum());
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
