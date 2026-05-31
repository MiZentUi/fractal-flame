package com.fractalflame.gateway.service;

import org.springframework.core.io.ByteArrayResource;
import org.springframework.core.io.Resource;
import org.springframework.stereotype.Service;

import com.fractalflame.gateway.mapper.UserMapper;
import com.fractalflame.gateway.model.AuthRequest;
import com.fractalflame.gateway.model.UserRequest;
import com.fractalflame.gateway.model.UserResponse;
import com.fractalflame.iam.proto.GetImageRequest;
import com.fractalflame.iam.proto.IAMServiceGrpc.IAMServiceBlockingStub;
import com.fractalflame.iam.proto.LoginResponse;
import com.fractalflame.iam.proto.RefreshRequest;
import com.fractalflame.iam.proto.RefreshResponse;
import com.fractalflame.iam.proto.RegisterResponse;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class IamService {
    private final IAMServiceBlockingStub stub;
    private final UserMapper mapper;

    public UserResponse getById(Long id) {
        return mapper.toUserResponse(stub.getUser(mapper.toGetUserRequest(id)).getUser());
    }

    public Resource getImageByName(String name) {
        return new ByteArrayResource(stub.getImage(GetImageRequest.newBuilder()
                .setName(name)
                .build()).getImage().toByteArray());
    }

    public LoginResponse login(AuthRequest authRequest) {
        return stub.login(com.fractalflame.iam.proto.AuthRequest.newBuilder()
                .setUsername(authRequest.getUsername())
                .setPassword(authRequest.getPassword())
                .build());
    }

    public UserResponse updateUser(UserRequest userRequest) {
        return mapper.toUserResponse(stub.updateUser(mapper.toUpdateUserRequest(userRequest)).getUser());
    }

    public RegisterResponse signUp(AuthRequest authRequest) {
        return stub.register(com.fractalflame.iam.proto.AuthRequest.newBuilder()
                .setUsername(authRequest.getUsername())
                .setPassword(authRequest.getPassword())
                .build());
    }

    public RefreshResponse refresh(String refreshToken) {
        return stub.refresh(RefreshRequest.newBuilder()
                .setRefreshToken(refreshToken)
                .build());
    }
}
