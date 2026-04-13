package academy.utils;

import academy.model.functions.DiscFunction;
import academy.model.functions.EyefishFunction;
import academy.model.functions.Function;
import academy.model.functions.HeartFunction;
import academy.model.functions.HorseshoeFunction;
import academy.model.functions.LinearFunction;
import academy.model.functions.PolarFunction;
import academy.model.functions.SinusoidalFunction;
import academy.model.functions.SphericalFunction;
import academy.model.functions.SwirlFunction;
import java.lang.reflect.InvocationTargetException;
import java.util.HashMap;
import java.util.Map;

public class FunctionBuilder {
    private static final Map<String, Class<? extends Function>> registry = new HashMap<>();

    static {
        registry.put("disc", DiscFunction.class);
        registry.put("eyefish", EyefishFunction.class);
        registry.put("heart", HeartFunction.class);
        registry.put("horseshoe", HorseshoeFunction.class);
        registry.put("linear", LinearFunction.class);
        registry.put("polar", PolarFunction.class);
        registry.put("sinusoidal", SinusoidalFunction.class);
        registry.put("spherical", SphericalFunction.class);
        registry.put("swirl", SwirlFunction.class);
    }

    private String name;
    private double weight;

    public FunctionBuilder byName(String name) {
        this.name = name;
        return this;
    }

    public FunctionBuilder withWeight(double width) {
        this.weight = width;
        return this;
    }

    public Function build() {
        try {
            return registry.get(name).getConstructor(Double.class).newInstance(weight);
        } catch (NoSuchMethodException
                | InstantiationException
                | IllegalAccessException
                | InvocationTargetException e) {
            throw new RuntimeException(e);
        }
    }
}
