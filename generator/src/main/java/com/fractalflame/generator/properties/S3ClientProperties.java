package com.fractalflame.generator.properties;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.validation.annotation.Validated;

import jakarta.validation.constraints.NotEmpty;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@ConfigurationProperties(prefix = "app.s3")
@Validated
@Setter
@Getter
@NoArgsConstructor
@AllArgsConstructor
public class S3ClientProperties {

    @NotEmpty
    private String accessKey;

    @NotEmpty
    private String secretKey;

    @NotEmpty
    private String url;
}
