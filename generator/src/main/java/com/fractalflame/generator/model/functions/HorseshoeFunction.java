package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

public class HorseshoeFunction extends FunctionModel {
    public HorseshoeFunction() {
        super(0.0);
    }

    public HorseshoeFunction(Double width) {
        super(width);
    }

    @Override
    protected Point f(Point point) {
        double r = Math.sqrt(point.x() * point.x() + point.y() * point.y());
        double newX = (point.x() - point.y()) * (point.x() + point.y()) / r;
        double newY = 2 * point.x() * point.y() / r;
        return new Point(newX, newY);
    }
}
