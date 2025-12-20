package academy.model;

import java.util.Random;

public class AffineTransformation {
    public record Params(double a, double b, double c, double d, double e, double f) {}

    private static Random random = new Random(System.currentTimeMillis());

    private final Params params;
    private final Pixel color;

    public AffineTransformation(Params params) {
        this.params = params;
        this.color = Pixel.genColor(random);
    }

    public Params getParams() {
        return params;
    }

    public Pixel getColor() {
        return new Pixel(color);
    }

    public Point transform(Point point) {
        double new_x = params.a() * point.x() + params.b * point.y() + params.c();
        double new_y = params.d() * point.x() + params.e() * point.y() + params.f();
        return new Point(new_x, new_y);
    }

    public static void setSeed(long seed) {
        random = new Random(seed);
    }
}
