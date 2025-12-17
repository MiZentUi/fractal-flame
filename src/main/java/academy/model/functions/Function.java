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
}
