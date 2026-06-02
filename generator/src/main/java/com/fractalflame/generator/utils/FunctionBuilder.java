package com.fractalflame.generator.utils;

import com.fractalflame.generator.exception.FunctionException;
import com.fractalflame.generator.model.functions.DiscFunction;
import com.fractalflame.generator.model.functions.EyefishFunction;
import com.fractalflame.generator.model.functions.FunctionModel;
import com.fractalflame.generator.model.functions.HeartFunction;
import com.fractalflame.generator.model.functions.HorseshoeFunction;
import com.fractalflame.generator.model.functions.LinearFunction;
import com.fractalflame.generator.model.functions.PolarFunction;
import com.fractalflame.generator.model.functions.SinusoidalFunction;
import com.fractalflame.generator.model.functions.SphericalFunction;
import com.fractalflame.generator.model.functions.SwirlFunction;
import java.lang.reflect.InvocationTargetException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.springframework.stereotype.Component;

@Component
public class FunctionBuilder {
    private static final Map<String, Class<? extends FunctionModel>> registry = new HashMap<>();

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

    private FunctionBuilder() {
    }

    public static List<String> getFunctionsNames() {
        return registry.keySet().stream().toList();
    }

    public static FunctionModel build(String name, Double weight) {
        if (!registry.containsKey(name)) {
            throw new FunctionException("Function with \"" + name + "\" name not found!");
        }
        try {
            return registry.get(name).getConstructor(Double.class).newInstance(weight);
        } catch (NoSuchMethodException
                | InstantiationException
                | IllegalAccessException
                | InvocationTargetException e) {
            throw new FunctionException(e);
        }
    }
}