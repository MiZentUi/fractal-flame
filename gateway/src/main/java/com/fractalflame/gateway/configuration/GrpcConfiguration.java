package com.fractalflame.gateway.configuration;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.grpc.client.GrpcChannelFactory;

import com.fractalflame.gateway.interceptor.AuthInterceptor;
import com.fractalflame.generator.proto.FractalsGrpc;
import com.fractalflame.generator.proto.ImagesGrpc;

@Configuration
public class GrpcConfiguration {

    @Bean
    FractalsGrpc.FractalsBlockingStub fractalsBlockingStub(GrpcChannelFactory channels,
            AuthInterceptor authInterceptor) {
        return FractalsGrpc.newBlockingStub(channels.createChannel("generator"));
    }

    @Bean
    FractalsGrpc.FractalsStub fractalsStub(GrpcChannelFactory channels, AuthInterceptor authInterceptor) {
        return FractalsGrpc.newStub(channels.createChannel("generator"));
    }

    @Bean
    ImagesGrpc.ImagesBlockingStub imagesStub(GrpcChannelFactory channels) {
        return ImagesGrpc.newBlockingStub(channels.createChannel("generator"));
    }
}
