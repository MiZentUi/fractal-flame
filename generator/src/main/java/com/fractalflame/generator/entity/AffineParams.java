package com.fractalflame.generator.entity;

import org.hibernate.validator.constraints.Length;

import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;
import jakarta.validation.constraints.NotNull;
import lombok.Getter;

@Entity
@Table(name = "affine_params")
@Getter
public class AffineParams {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "fractal_id")
    private Fractal fractal;

    @Length(min = 7, max = 7)
    private String color;

    @NotNull
    private Double a;

    @NotNull
    private Double b;

    @NotNull
    private Double c;

    @NotNull
    private Double d;

    @NotNull
    private Double e;

    @NotNull
    private Double f;
}
