package com.fractalflame.generator.mapper;

import java.util.List;

import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.MappingConstants;
import org.mapstruct.Named;
import org.mapstruct.ReportingPolicy;

import com.fractalflame.generator.entity.AffineParams;
import com.fractalflame.generator.model.AffineTransformation;
import com.fractalflame.generator.model.Pixel;

@Mapper(componentModel = MappingConstants.ComponentModel.SPRING, unmappedTargetPolicy = ReportingPolicy.IGNORE)
public interface AffineMapper {

    @Mapping(target = "color", source = "color", qualifiedByName = "colorToPixel")
    @Mapping(target = "params.a", source = "a")
    @Mapping(target = "params.b", source = "b")
    @Mapping(target = "params.c", source = "c")
    @Mapping(target = "params.d", source = "d")
    @Mapping(target = "params.e", source = "e")
    @Mapping(target = "params.f", source = "f")
    AffineTransformation toAffineTransformation(AffineParams affineParams);

    List<AffineTransformation> toAffineTransformationList(List<AffineParams> affineParams);

    com.fractalflame.generator.proto.AffineParams toProtoAffineParams(AffineParams affineParams);

    List<com.fractalflame.generator.proto.AffineParams> toProtoAffineParamsList(List<AffineParams> affineParams);

    AffineParams fromProtoAffineParams(com.fractalflame.generator.proto.AffineParams affineParams);

    @Named("colorToPixel")
    default Pixel colorToPixel(String hexColor) {
        int hex = Integer.parseInt(hexColor.substring(1), 16);
        int r = (hex >> 16) & 0xFF;
        int g = (hex >> 8) & 0xFF;
        int b = hex & 0xFF;
        return new Pixel(r, g, b);
    }
}
