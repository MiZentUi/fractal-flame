package com.fractalflame.gateway.interceptor;

import org.springframework.grpc.client.GlobalClientInterceptor;
import org.springframework.stereotype.Component;

import com.fractalflame.gateway.filter.JwtAuthFilter;
import com.fractalflame.gateway.service.JwtService;

import io.grpc.CallOptions;
import io.grpc.Channel;
import io.grpc.ClientCall;
import io.grpc.ClientInterceptor;
import io.grpc.ForwardingClientCall;
import io.grpc.Metadata;
import io.grpc.MethodDescriptor;
import lombok.RequiredArgsConstructor;

@Component
@GlobalClientInterceptor
@RequiredArgsConstructor
public class AuthInterceptor implements ClientInterceptor {
    private final JwtService jwtService;

    @Override
    public <ReqT, RespT> ClientCall<ReqT, RespT> interceptCall(MethodDescriptor<ReqT, RespT> method,
            CallOptions callOptions, Channel next) {
        return new ForwardingClientCall.SimpleForwardingClientCall<ReqT, RespT>(next.newCall(method, callOptions)) {
            @Override
            public void start(Listener<RespT> responseListener, Metadata headers) {
                var token = jwtService.getToken();
                if (token != null && !token.isEmpty() && jwtService.validateToken(token)) {
                    headers.put(Metadata.Key.of(JwtAuthFilter.HEADER_NAME, Metadata.ASCII_STRING_MARSHALLER),
                            JwtAuthFilter.BEARER_PREFIX + token);
                }
                super.start(responseListener, headers);
            }
        };
    }
}
