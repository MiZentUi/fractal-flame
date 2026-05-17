package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

public class SwirlFunction extends FunctionModel {
    public SwirlFunction() {
        super(0.0);
    }

    public SwirlFunction(Double width) {
        super(width);
    }

    @Override
    protected Point f(Point point) {
        double rSq = point.x() * point.x() + point.y() * point.y();
        double newX = point.x() * Math.sin(rSq) - point.y() * Math.cos(rSq);
        double newY = point.x() * Math.cos(rSq) + point.y() * Math.sin(rSq);
        return new Point(newX, newY);
    }
}
