package academy.model.functions;

import academy.model.Point;

public class PolarFunction extends Function {
    public PolarFunction() {
        super(0.0);
    }

    public PolarFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point F(Point point) {
        double r = Math.sqrt(point.x() * point.x() + point.y() * point.y());
        double new_x = Math.atan(point.y() / point.x()) / Math.PI;
        double new_y = r - 1;
        return new Point(new_x, new_y);
    }
}
