package academy.model.functions;

import academy.model.Point;

public class SwirlFunction extends Function {
    public SwirlFunction() {
        super(0.0);
    }

    public SwirlFunction(Double width) {
        super(width);
    }

    @Override
    protected Point F(Point point) {
        double r_sq = point.x() * point.x() + point.y() * point.y();
        double new_x = point.x() * Math.sin(r_sq) - point.y() * Math.cos(r_sq);
        double new_y = point.x() * Math.cos(r_sq) + point.y() * Math.sin(r_sq);
        return new Point(new_x, new_y);
    }
}
