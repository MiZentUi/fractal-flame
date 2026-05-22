package com.fractalflame.gateway.mapper;

import org.mapstruct.Mapper;
import org.mapstruct.MappingConstants;
import org.mapstruct.CollectionMappingStrategy;
import org.mapstruct.NullValueCheckStrategy;
import org.mapstruct.ReportingPolicy;

import com.fractalflame.gateway.model.AffineParams;

@Mapper(componentModel = MappingConstants.ComponentModel.SPRING, collectionMappingStrategy = CollectionMappingStrategy.ADDER_PREFERRED, nullValueCheckStrategy = NullValueCheckStrategy.ALWAYS, unmappedTargetPolicy = ReportingPolicy.IGNORE)
public interface AffineParamsMapper {
    AffineParams toAffineParams(com.fractalflame.generator.proto.AffineParams params);

    com.fractalflame.generator.proto.AffineParams fromAffineParams(AffineParams params);
}
