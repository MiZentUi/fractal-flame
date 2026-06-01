package com.fractalflame.gateway.controller;

import java.time.Duration;

import org.springframework.core.io.Resource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseCookie;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.fractalflame.gateway.api.UsersApi;
import com.fractalflame.gateway.model.AccessToken;
import com.fractalflame.gateway.model.AuthRequest;
import com.fractalflame.gateway.model.UserRequest;
import com.fractalflame.gateway.model.UserResponse;
import com.fractalflame.gateway.service.IamService;

import jakarta.validation.Valid;
import jakarta.validation.constraints.NotNull;
import lombok.RequiredArgsConstructor;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class UsersController implements UsersApi {
    private final IamService service;

    @Override
    public ResponseEntity<UserResponse> getUser(Long id) {
        return ResponseEntity.ok(service.getById(id));
    }

    @Override
    public ResponseEntity<Resource> getUserImage(String name) {
        return ResponseEntity.ok(service.getImageByName(name));
    }

    @Override
    public ResponseEntity<AccessToken> login(@Valid AuthRequest authRequest) {
        var response = service.login(authRequest);

        var refreshCookie = createRefreshCookie(response.getRefreshToken());

        return ResponseEntity.ok()
                .header(HttpHeaders.SET_COOKIE, refreshCookie.toString())
                .body(AccessToken.builder()
                        .accessToken(response.getAccessToken())
                        .build());
    }

    @Override
    public ResponseEntity<UserResponse> patchUser(@Valid UserRequest userRequest) {
        return ResponseEntity.ok(service.updateUser(userRequest));
    }

    @Override
    public ResponseEntity<AccessToken> refresh(@NotNull String refreshToken) {
        var response = service.refresh(refreshToken);

        var refreshCookie = createRefreshCookie(response.getRefreshToken());

        return ResponseEntity.ok()
                .header(HttpHeaders.SET_COOKIE, refreshCookie.toString())
                .body(AccessToken.builder()
                        .accessToken(response.getAccessToken())
                        .build());
    }

    @Override
    public ResponseEntity<Void> signUp(@Valid AuthRequest authRequest) {
        service.signUp(authRequest);
        return ResponseEntity.ok().build();
    }

    private ResponseCookie createRefreshCookie(String refreshToken) {
        return ResponseCookie.from("refresh_token", refreshToken)
                .httpOnly(false) // a forced measure for raw logout
                .secure(true)
                .sameSite("Strict")
                .path("/api/v1/refresh")
                .maxAge(Duration.ofDays(7))
                .build();
    }
}
