package com.fractalflame.generator.model;

import java.util.Random;
import java.util.concurrent.atomic.LongAdder;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
public class Pixel {

    @Getter(AccessLevel.NONE)
    private final LongAdder red = new LongAdder();

    @Getter(AccessLevel.NONE)
    private final LongAdder green = new LongAdder();

    @Getter(AccessLevel.NONE)
    private final LongAdder blue = new LongAdder();

    @Getter(AccessLevel.NONE)
    private final LongAdder alpha = new LongAdder();

    private LongAdder counter = new LongAdder();

    private Double normal = 0.0;

    public Pixel(Pixel pixel) {
        if (pixel != null) {
            this.red.add(pixel.red.sum());
            this.green.add(pixel.green.sum());
            this.blue.add(pixel.blue.sum());
            this.alpha.add(pixel.alpha.sum());
        }
    }

    public Pixel(int red, int green, int blue) {
        this.red.add(red);
        this.green.add(green);
        this.blue.add(blue);
        alpha.add(255);
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
        int iAlpha = Math.toIntExact(alpha.sum());
        int iRed = Math.toIntExact(red.sum());
        int iGreen = Math.toIntExact(green.sum());
        int iBlue = Math.toIntExact(blue.sum());
        return (iAlpha << 24) | (iRed << 16) | (iGreen << 8) | iBlue;
    }

    public static Pixel genColor(Random random) {
        int low = 0;
        int high = 255;
        int red = random.nextInt(low, high);
        int green = random.nextInt(low, high);
        int blue = random.nextInt(low, high);
        return new Pixel(red, green, blue);
    }
}
