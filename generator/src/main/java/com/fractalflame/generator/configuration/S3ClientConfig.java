package com.fractalflame.generator.configuration;

import java.net.URI;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.fractalflame.generator.properties.S3ClientProperties;

import lombok.RequiredArgsConstructor;
import software.amazon.awssdk.auth.credentials.AwsBasicCredentials;
import software.amazon.awssdk.auth.credentials.StaticCredentialsProvider;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.s3.S3Client;
import software.amazon.awssdk.services.s3.S3Configuration;

@Configuration
@RequiredArgsConstructor
public class S3ClientConfig {
    private final S3ClientProperties properties;

    @Bean
    S3Client s3Client() {
        AwsBasicCredentials credentials = AwsBasicCredentials.create(properties.getAccessKey(),
                properties.getSecretKey());

        return S3Client.builder()
                .endpointOverride(URI.create(properties.getUrl()))
                .credentialsProvider(StaticCredentialsProvider.create(credentials))
                .region(Region.EU_CENTRAL_1)
                .serviceConfiguration(S3Configuration.builder()
                        .pathStyleAccessEnabled(true)
                        .build())
                .build();
    }
}
