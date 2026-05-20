package com.fractalflame.generator.mapper;

import java.time.Instant;
import java.util.List;

import org.mapstruct.CollectionMappingStrategy;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.MappingConstants;
import org.mapstruct.Named;
import org.mapstruct.NullValueCheckStrategy;
import org.mapstruct.ReportingPolicy;

import com.fractalflame.generator.entity.Fractal;
import com.fractalflame.generator.proto.FractalRequest;
import com.fractalflame.generator.proto.FractalResponse;
import com.google.protobuf.Timestamp;

@Mapper(componentModel = MappingConstants.ComponentModel.SPRING, uses = { AffineMapper.class,
        FunctionMapper.class }, collectionMappingStrategy = CollectionMappingStrategy.ADDER_PREFERRED, nullValueCheckStrategy = NullValueCheckStrategy.ALWAYS, unmappedTargetPolicy = ReportingPolicy.IGNORE)
public interface FractalMapper {
    @Mapping(target = "functionsList", source = "functions")
    @Mapping(target = "affineParamsList", source = "affineParams")
    @Mapping(target = "created", source = "created", qualifiedByName = "toTimestamp")
    FractalResponse toFractalResponse(Fractal fractal);

    List<FractalResponse> toFractalResponseList(List<Fractal> fractals);

    @Mapping(target = "functions", source = "functionsList")
    @Mapping(target = "affineParams", source = "affineParamsList")
    @Mapping(target = "created", ignore = true)
    Fractal fromFractalRequest(FractalRequest fractalRequest);

    @Named("toTimestamp")
    default Timestamp toTimestamp(Instant instant) {
        return Timestamp.newBuilder()
                .setSeconds(instant.getEpochSecond())
                .setNanos(instant.getNano())
                .build();
    }
}
