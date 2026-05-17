package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

public class EyefishFunction extends FunctionModel {
    public EyefishFunction() {
        super(0.0);
    }

    public EyefishFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point f(Point point) {
        double r = Math.sqrt(point.x() * point.x() + point.y() * point.y());
        return new Point(point).multiply(2 / (r - 1));
    }
}
