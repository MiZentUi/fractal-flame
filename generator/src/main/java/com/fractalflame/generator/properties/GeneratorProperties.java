package com.fractalflame.generator.properties;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.validation.annotation.Validated;

import jakarta.validation.constraints.NotNull;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@ConfigurationProperties
@Validated
@Setter
@Getter
@NoArgsConstructor
public class GeneratorProperties {

    @NotNull
    private Integer count;

    @NotNull
    private Integer threads;
}
