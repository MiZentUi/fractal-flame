package com.fractalflame.generator.model;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@RequiredArgsConstructor
@Getter
public class AffineTransformation {
    public record Params(double a, double b, double c, double d, double e, double f) {
    }

    private final Params params;
    private final Pixel color;

    public Pixel getColor() {
        return new Pixel(color);
    }

    public Point transform(Point point) {
        double newX = params.a() * point.x() + params.b * point.y() + params.c();
        double newY = params.d() * point.x() + params.e() * point.y() + params.f();
        return new Point(newX, newY);
    }
}
