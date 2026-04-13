package academy.model.functions;

import academy.model.Point;

public abstract class Function {
    double width;

    public Function(Double weight) {
        this.width = weight;
    }

    public Point transform(Point point) {
        return F(point);
    }

    public double getWidth() {
        return width;
    }

    protected abstract Point F(Point point);

    @Override
    public boolean equals(Object o) {
        if (o == null || getClass() != o.getClass()) return false;

        Function function = (Function) o;
        return Double.compare(width, function.width) == 0;
    }

    @Override
    public int hashCode() {
        return Double.hashCode(width);
    }
}
