package com.fractalflame.generator;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.List;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.annotation.Import;

import com.fractalflame.generator.entity.AffineParams;
import com.fractalflame.generator.entity.Fractal;
import com.fractalflame.generator.entity.Function;
import com.fractalflame.generator.mapper.AffineMapper;
import com.fractalflame.generator.mapper.FractalMapper;
import com.fractalflame.generator.mapper.FunctionMapper;
import com.fractalflame.generator.proto.FractalRequest;

@Import(TestcontainersConfiguration.class)
@SpringBootTest
class MappersTests {

    @Autowired
    private AffineMapper affineMapper;

    @Autowired
    private FunctionMapper functionMapper;

    @Autowired
    private FractalMapper fractalMapper;

    @Test
    void toAffineTransformation() {
        var affineParams = AffineParams.builder()
                .color("#ffffff")
                .a(1.0)
                .b(2.0)
                .c(3.0)
                .d(4.0)
                .e(5.0)
                .f(6.0)
                .build();

        var affineTransformation = affineMapper.toAffineTransformation(affineParams);

        var params = affineTransformation.getParams();
        assertEquals(affineMapper.colorToPixel(affineParams.getColor()).toRGB(),
                affineTransformation.getColor().toRGB());
        assertEquals(affineParams.getA(), params.a());
        assertEquals(affineParams.getB(), params.b());
        assertEquals(affineParams.getC(), params.c());
        assertEquals(affineParams.getD(), params.d());
        assertEquals(affineParams.getE(), params.e());
        assertEquals(affineParams.getF(), params.f());
    }

    @Test
    void fromProtoAffineParams() {
        var protoAffineParams = com.fractalflame.generator.proto.AffineParams.newBuilder()
                .setColor("#ffffff")
                .setA(1)
                .setB(2)
                .setC(3)
                .setD(4)
                .setE(5)
                .setF(6)
                .build();

        var affineParams = affineMapper.fromProtoAffineParams(protoAffineParams);

        assertEquals(protoAffineParams.getColor(), affineParams.getColor());
        assertEquals(protoAffineParams.getA(), affineParams.getA());
        assertEquals(protoAffineParams.getB(), affineParams.getB());
        assertEquals(protoAffineParams.getC(), affineParams.getC());
        assertEquals(protoAffineParams.getD(), affineParams.getD());
        assertEquals(protoAffineParams.getE(), affineParams.getE());
        assertEquals(protoAffineParams.getF(), affineParams.getF());
    }

    @Test
    void toFunctionModel() {
        var function = Function.builder()
                .name("disc")
                .weight(1.5)
                .build();

        var functionModel = functionMapper.toFunctionModel(function);

        assertEquals(function.getName(), functionModel.getName());
        assertEquals(function.getWeight(), functionModel.getWeight());
    }

    @Test
    void fromProtoFunction() {
        var protoFunction = com.fractalflame.generator.proto.Function.newBuilder()
                .setName("func")
                .setWeight(1.5)
                .build();

        var function = functionMapper.fromProtoFunction(protoFunction);

        assertEquals(protoFunction.getName(), function.getName());
        assertEquals(protoFunction.getWeight(), function.getWeight());
    }

    @Test
    void toFractalResponse() {
        var function = Function.builder()
                .name("disc")
                .weight(1.5)
                .build();

        var affineParams = AffineParams.builder()
                .color("#ffffff")
                .a(1.0)
                .b(2.0)
                .c(3.0)
                .d(4.0)
                .e(5.0)
                .f(6.0)
                .build();

        var fractal = Fractal.builder()
                .width(1)
                .height(2)
                .functions(List.of(function))
                .affineParams(List.of(affineParams))
                .build();

        var fractalResponse = fractalMapper.toFractalResponse(fractal);

        assertEquals(fractal.getWidth(), fractalResponse.getWidth());
        assertEquals(fractal.getHeight(), fractalResponse.getHeight());
    }

    @Test
    void fromFractalRequest() {
        var protoFunction = com.fractalflame.generator.proto.Function.newBuilder()
                .setName("func")
                .setWeight(1.5)
                .build();

        var protoAffineParams = com.fractalflame.generator.proto.AffineParams.newBuilder()
                .setColor("#ffffff")
                .setA(1)
                .setB(2)
                .setC(3)
                .setD(4)
                .setE(5)
                .setF(6)
                .build();

        var fractalRequest = FractalRequest.newBuilder()
                .setWidth(1)
                .setHeight(2)
                .addAllFunctions(List.of(protoFunction))
                .addAllAffineParams(List.of(protoAffineParams))
                .build();

        var fractal = fractalMapper.fromFractalRequest(fractalRequest);

        assertEquals(fractalRequest.getWidth(), fractal.getWidth());
        assertEquals(fractalRequest.getHeight(), fractal.getHeight());
        assertEquals(fractalRequest.getFunctionsList().size(), fractal.getFunctions().size());
        assertEquals(fractalRequest.getAffineParamsList().size(), fractal.getAffineParams().size());
    }
}
