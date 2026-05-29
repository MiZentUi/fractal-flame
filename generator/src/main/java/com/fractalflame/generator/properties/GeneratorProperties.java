package com.fractalflame.generator.properties;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.validation.annotation.Validated;

import jakarta.validation.constraints.Min;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@ConfigurationProperties(prefix = "app.generator")
@Validated
@Setter
@Getter
@NoArgsConstructor
@AllArgsConstructor
public class GeneratorProperties {

    @Min(1)
    private Integer count;

    @Min(1)
    private Integer threads;

    @Min(1)
    private Integer maxWidth;

    @Min(1)
    private Integer maxHeight;

    @Min(1)
    private Integer maxIterations;
}
