package com.fractalflame.gateway.mapper;

import java.time.Instant;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.List;

import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.MappingConstants;
import org.mapstruct.Named;
import org.mapstruct.CollectionMappingStrategy;
import org.mapstruct.NullValueCheckStrategy;
import org.mapstruct.ReportingPolicy;

import com.fractalflame.gateway.model.FractalRequest;
import com.fractalflame.gateway.model.FractalResponse;
import com.fractalflame.generator.proto.FractalsRequest;
import com.fractalflame.generator.proto.Order;
import com.google.protobuf.Timestamp;

@Mapper(componentModel = MappingConstants.ComponentModel.SPRING, uses = { FunctionMapper.class,
        AffineParamsMapper.class }, collectionMappingStrategy = CollectionMappingStrategy.ADDER_PREFERRED, nullValueCheckStrategy = NullValueCheckStrategy.ALWAYS, unmappedTargetPolicy = ReportingPolicy.IGNORE)
public interface FractalMapper {

    @Mapping(target = "order", source = "order", qualifiedByName = "orderEnum")
    FractalsRequest toFractalsRequest(Integer page, Integer count, String sort, String order, Long userId);

    @Mapping(target = "functions", source = "functionsList")
    @Mapping(target = "affineParams", source = "affineParamsList")
    @Mapping(target = "created", source = "created", qualifiedByName = "fromTimestamp")
    FractalResponse toFractalResponse(com.fractalflame.generator.proto.FractalResponse fractalResponse);

    List<FractalResponse> toFractalResponseList(List<com.fractalflame.generator.proto.FractalResponse> fractalResponse);

    @Mapping(target = "functionsList", source = "functions")
    @Mapping(target = "affineParamsList", source = "affineParams")
    com.fractalflame.generator.proto.FractalRequest fromFractalRequest(FractalRequest fractalRequest);

    @Named("fromTimestamp")
    default OffsetDateTime fromTimestamp(Timestamp timestamp) {
        return Instant.ofEpochSecond(timestamp.getSeconds(), timestamp.getNanos()).atOffset(ZoneOffset.UTC);
    }

    @Named("orderEnum")
    default Order fromTimestamp(String order) {
        if (order == null) {
            return null;
        }
        return Order.valueOf(order.toUpperCase());
    }
}
