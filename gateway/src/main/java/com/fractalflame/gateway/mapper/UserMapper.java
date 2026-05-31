package com.fractalflame.gateway.mapper;

import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.MappingConstants;
import org.mapstruct.Named;

import java.util.List;

import org.mapstruct.CollectionMappingStrategy;
import org.mapstruct.NullValueCheckStrategy;
import org.mapstruct.ReportingPolicy;

import com.fractalflame.gateway.model.UserRequest;
import com.fractalflame.gateway.model.UserResponse;
import com.fractalflame.iam.proto.GetUserRequest;
import com.fractalflame.iam.proto.UpdateUserRequest;
import com.fractalflame.iam.proto.User;
import com.google.protobuf.StringValue;

@Mapper(componentModel = MappingConstants.ComponentModel.SPRING, collectionMappingStrategy = CollectionMappingStrategy.ADDER_PREFERRED, nullValueCheckStrategy = NullValueCheckStrategy.ALWAYS, unmappedTargetPolicy = ReportingPolicy.IGNORE)
public interface UserMapper {

    @Mapping(target = "id", source = "userId")
    UserResponse toUserResponse(User user);

    List<UserResponse> toUserResponseList(List<User> user);

    GetUserRequest toGetUserRequest(Long userId);

    @Mapping(target = "username", source = "username", qualifiedByName = "stringValue")
    @Mapping(target = "password", source = "password", qualifiedByName = "stringValue")
    @Mapping(target = "image", source = "image", qualifiedByName = "stringValue")
    UpdateUserRequest toUpdateUserRequest(UserRequest userRequest);

    @Named("stringValue")
    default StringValue toStringValue(String value) {
        if (value == null) {
            return null;
        }
        return StringValue.of(value);
    }
}
