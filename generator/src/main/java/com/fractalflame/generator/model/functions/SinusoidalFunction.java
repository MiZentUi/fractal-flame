package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

public class SinusoidalFunction extends FunctionModel {
    public SinusoidalFunction() {
        super(0.0);
    }

    public SinusoidalFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point f(Point point) {
        double newX = Math.sin(point.x());
        double newY = Math.sin(point.y());
        return new Point(newX, newY);
    }
}
