package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

public class HeartFunction extends FunctionModel {
    public HeartFunction() {
        super(0.0);
    }

    public HeartFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point f(Point point) {
        double r = Math.sqrt(point.x() * point.x() + point.y() * point.y());
        double newX = r * Math.sin(r * Math.atan(point.y() / point.x()));
        double newY = -1 * r * Math.cos(r * Math.atan(point.y() / point.x()));
        return new Point(newX, newY);
    }
}
