package com.fractalflame.generator.mapper;

import java.util.List;
import org.springframework.stereotype.Component;

import com.fractalflame.generator.entity.Function;
import com.fractalflame.generator.model.functions.FunctionModel;
import com.fractalflame.generator.utils.FunctionBuilder;

import lombok.RequiredArgsConstructor;

@Component
@RequiredArgsConstructor
public class FunctionMapper {
    public FunctionModel toFunctionModel(Function function) {
        return FunctionBuilder.build(function.getName(), function.getWeight());
    }

    public List<FunctionModel> toFunctionModelList(List<Function> functions) {
        return functions.stream().map(this::toFunctionModel).toList();
    }

    public com.fractalflame.generator.proto.Function toProtoFunction(Function function) {
        return com.fractalflame.generator.proto.Function.newBuilder()
                .setName(function.getName())
                .setWeight(function.getWeight())
                .build();
    }

    public List<com.fractalflame.generator.proto.Function> toProtoFunctionList(List<Function> functions) {
        return functions.stream().map(this::toProtoFunction).toList();
    }

    public Function fromProtoFunction(com.fractalflame.generator.proto.Function function) {
        return Function.builder()
                .name(function.getName())
                .weight(function.getWeight())
                .build();
    }
}
