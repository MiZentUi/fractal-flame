package academy.model.functions;

import academy.model.Point;

public class LinearFunction extends Function {
    public LinearFunction() {
        super(0.0);
    }

    public LinearFunction(Double weight) {
        super(weight);
    }

    @Override
    protected Point F(Point point) {
        return point;
    }
}
