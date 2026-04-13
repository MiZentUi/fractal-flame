package academy.model.functions;

import academy.model.Point;

public class HeartFunction extends Function {
    public HeartFunction() {
        super(0.0);
    }

    public HeartFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point F(Point point) {
        double r = Math.sqrt(point.x() * point.x() + point.y() * point.y());
        double new_x = r * Math.sin(r * Math.atan((double) point.y() / point.x()));
        double new_y = -1 * r * Math.cos(r * Math.atan((double) point.y() / point.x()));
        return new Point(new_x, new_y);
    }
}
