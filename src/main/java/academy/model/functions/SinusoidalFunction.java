package academy.model.functions;

import academy.model.Point;

public class SinusoidalFunction extends Function {
    public SinusoidalFunction() {
        super(0.0);
    }

    public SinusoidalFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point F(Point point) {
        double new_x = Math.sin(point.x);
        double new_y = Math.sin(point.y);
        return new Point(new_x, new_y);
    }
}
