package academy.model.functions;

import academy.model.Point;

public class EyefishFunction extends Function {
    public EyefishFunction() {
        super(0.0);
    }

    public EyefishFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point F(Point point) {
        double r = Math.sqrt(point.x() * point.x() + point.y() * point.y());
        return new Point(point).multiply(2 / (r - 1));
    }
}
