package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

public class LinearFunction extends FunctionModel {
    public LinearFunction() {
        super(0.0);
    }

    public LinearFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point f(Point point) {
        return point;
    }

    @Override
    public String getName() {
        return "linear";
    }
}
