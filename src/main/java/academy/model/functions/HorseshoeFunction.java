package academy.model.functions;

import academy.model.AffineTransformation;
import academy.model.Point;

public class HorseshoeFunction extends Function {
    public HorseshoeFunction() {
        super(0.0);
    }

    public HorseshoeFunction(Double width) {
        super(width);
    }

    @Override
    protected Point F(Point point) {
        double r = Math.sqrt(point.x * point.x + point.y * point.y);
        double new_x = (point.x - point.y) * (point.x + point.y) / r;
        double new_y = 2 * point.x * point.y / r;
        return new Point(new_x, new_y);
    }
}
