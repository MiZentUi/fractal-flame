package academy.model.functions;

import academy.model.Point;

public class SphericalFunction extends Function {
    public SphericalFunction() {
        super(0.0);
    }

    public SphericalFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point F(Point point) {
        double r_sq = point.x * point.x + point.y * point.y;
        double new_x = point.x / r_sq;
        double new_y = point.y / r_sq;
        return new Point(new_x, new_y);
    }
}
