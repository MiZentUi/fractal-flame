package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

public class DiscFunction extends FunctionModel {
    public DiscFunction() {
        super(0.0);
    }

    public DiscFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point f(Point point) {
        double a = Math.atan(point.y() / point.x());
        double r = Math.sqrt(point.x() * point.x() + point.y() * point.y());
        double newX = a * Math.sin(Math.PI * r) / Math.PI;
        double newY = a * Math.cos(Math.PI * r) / Math.PI;
        return new Point(newX, newY);
    }
}
