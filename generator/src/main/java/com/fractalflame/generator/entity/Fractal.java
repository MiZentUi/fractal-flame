package com.fractalflame.generator.entity;

import java.time.Instant;
import java.util.List;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.OneToMany;
import jakarta.persistence.Table;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.PositiveOrZero;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Entity
@Table(name = "fractals")
@Builder
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class Fractal {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "user_id")
    private Long userId;

    @Builder.Default
    private Instant created = Instant.now();

    @Min(1)
    private Integer width;

    @Min(1)
    private Integer height;

    @Min(1)
    private Integer iterationCount;

    @Min(1)
    private Integer symmetryLevel;

    @PositiveOrZero
    private Double gamma;

    private String image;

    @OneToMany(mappedBy = "fractal")
    private List<Function> functions;

    @OneToMany(mappedBy = "fractal")
    private List<AffineParams> affineParams;
}
