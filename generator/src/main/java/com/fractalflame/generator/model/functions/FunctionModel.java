package com.fractalflame.generator.model.functions;

import com.fractalflame.generator.model.Point;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public abstract class FunctionModel {
    double weight;

    public Point transform(Point point) {
        return f(point);
    }

    public void setWeight(double weight) {
        this.weight = weight;
    }

    public double getWeight() {
        return weight;
    }

    public abstract String getName();

    protected abstract Point f(Point point);

    @Override
    public boolean equals(Object o) {
        if (o == null || getClass() != o.getClass())
            return false;

        FunctionModel function = (FunctionModel) o;
        return Double.compare(weight, function.weight) == 0;
    }

    @Override
    public int hashCode() {
        return Double.hashCode(weight);
    }
}
