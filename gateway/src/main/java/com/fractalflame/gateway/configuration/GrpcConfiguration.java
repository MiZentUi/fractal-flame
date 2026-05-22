package com.fractalflame.gateway.configuration;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.grpc.client.GrpcChannelFactory;

import com.fractalflame.generator.proto.FractalsGrpc;
import com.fractalflame.generator.proto.ImagesGrpc;

@Configuration
public class GrpcConfiguration {

    @Bean
    FractalsGrpc.FractalsBlockingStub fractalsStub(GrpcChannelFactory channels) {
        return FractalsGrpc.newBlockingStub(channels.createChannel("local"));
    }

    @Bean
    ImagesGrpc.ImagesBlockingStub imagesStub(GrpcChannelFactory channels) {
        return ImagesGrpc.newBlockingStub(channels.createChannel("local"));
    }
}
