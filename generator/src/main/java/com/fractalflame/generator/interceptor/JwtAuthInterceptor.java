package com.fractalflame.generator.interceptor;

import java.util.List;

import org.springframework.grpc.server.GlobalServerInterceptor;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Component;

import com.fractalflame.generator.model.User;
import com.fractalflame.generator.service.JwtService;

import io.grpc.ForwardingServerCallListener;
import io.grpc.Metadata;
import io.grpc.ServerCall;
import io.grpc.ServerCall.Listener;
import lombok.RequiredArgsConstructor;
import io.grpc.ServerCallHandler;
import io.grpc.ServerInterceptor;
import io.grpc.Status;

@Component
@GlobalServerInterceptor
@RequiredArgsConstructor
public class JwtAuthInterceptor implements ServerInterceptor {
    private final JwtService jwtService;

    @Override
    public <ReqT, RespT> Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata headers,
            ServerCallHandler<ReqT, RespT> next) {

        Metadata.Key<String> authHeaderKey = Metadata.Key.of("Authorization", Metadata.ASCII_STRING_MARSHALLER);

        String authHeader = headers.get(authHeaderKey);

        if (authHeader == null || !authHeader.startsWith("Bearer ")) {
            return next.startCall(call, headers);
        }

        try {

            String token = authHeader.substring(7);

            jwtService.validateToken(token);

            var claims = jwtService.extractClaims(token);

            Long userId = Long.valueOf(claims.getSubject());

            String username = claims.get("name", String.class);

            List<SimpleGrantedAuthority> authorities = List.of(
                    new SimpleGrantedAuthority("ROLE_USER"));

            var principal = new User(
                    userId,
                    username,
                    authorities);

            UsernamePasswordAuthenticationToken authentication = new UsernamePasswordAuthenticationToken(
                    principal,
                    null,
                    authorities);

            SecurityContextHolder.getContext()
                    .setAuthentication(authentication);

            return new ForwardingServerCallListener.SimpleForwardingServerCallListener<>(
                    next.startCall(call, headers)) {

                @Override
                public void onComplete() {
                    SecurityContextHolder.clearContext();
                    super.onComplete();
                }

                @Override
                public void onCancel() {
                    SecurityContextHolder.clearContext();
                    super.onCancel();
                }
            };

        } catch (Exception e) {

            call.close(
                    Status.UNAUTHENTICATED
                            .withDescription("Invalid JWT"),
                    new Metadata());

            return new ServerCall.Listener<>() {
            };
        }
    }
}
