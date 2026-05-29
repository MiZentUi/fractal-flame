package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

public class PolarFunction extends FunctionModel {
    public PolarFunction() {
        super(0.0);
    }

    public PolarFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point f(Point point) {
        double r = Math.sqrt(point.x() * point.x() + point.y() * point.y());
        double newX = Math.atan(point.y() / point.x()) / Math.PI;
        double newY = r - 1;
        return new Point(newX, newY);
    }

    @Override
    public String getName() {
        return "polar";
    }
}
