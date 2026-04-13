package academy.model.functions;

import academy.model.Point;

public class DiscFunction extends Function {
    public DiscFunction() {
        super(0.0);
    }

    public DiscFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point F(Point point) {
        double a = Math.atan(point.y() / point.x());
        double r = Math.sqrt(point.x() * point.x() + point.y() * point.y());
        double new_x = a * Math.sin(Math.PI * r) / Math.PI;
        double new_y = a * Math.cos(Math.PI * r) / Math.PI;
        return new Point(new_x, new_y);
    }
}
