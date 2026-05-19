package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

public class SphericalFunction extends FunctionModel {
    public SphericalFunction() {
        super(0.0);
    }

    public SphericalFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point f(Point point) {
        double rSq = point.x() * point.x() + point.y() * point.y();
        double newX = point.x() / rSq;
        double newY = point.y() / rSq;
        return new Point(newX, newY);
    }

    @Override
    public String getName() {
        return "spherical";
    }
}
