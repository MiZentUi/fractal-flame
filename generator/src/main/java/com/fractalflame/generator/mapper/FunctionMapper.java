package com.fractalflame.generator.mapper;

import java.util.List;
import java.util.Map;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Component;

import com.fractalflame.generator.entity.Function;
import com.fractalflame.generator.model.functions.FunctionModel;

import lombok.RequiredArgsConstructor;

@Component
@RequiredArgsConstructor
public class FunctionMapper {

    @Qualifier("functionsByName")
    private final Map<String, FunctionModel> functions;

    public FunctionModel toFunctionModel(Function function) {
        var functionModel = functions.get(function.getName());
        functionModel.setWidth(function.getWeight());
        return functionModel;
    }

    public List<FunctionModel> toFunctionModelList(List<Function> functions) {
        return functions.stream().map(f -> toFunctionModel(f)).toList();
    }
}
