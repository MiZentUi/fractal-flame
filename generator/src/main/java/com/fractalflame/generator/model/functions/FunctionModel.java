package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

public abstract class FunctionModel {
    double width;

    protected FunctionModel(Double weight) {
        this.width = weight;
    }

    public Point transform(Point point) {
        return f(point);
    }

    public void setWidth(double width) {
        this.width = width;
    }

    public double getWidth() {
        return width;
    }

    protected abstract Point f(Point point);

    @Override
    public boolean equals(Object o) {
        if (o == null || getClass() != o.getClass())
            return false;

        FunctionModel function = (FunctionModel) o;
        return Double.compare(width, function.width) == 0;
    }

    @Override
    public int hashCode() {
        return Double.hashCode(width);
    }
}
