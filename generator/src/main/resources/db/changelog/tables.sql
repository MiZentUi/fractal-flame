--liquibase formatted sql

--changeset mizentui:tables

CREATE TABLE fractals
(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT,
    width INTEGER NOT NULL,
    height INTEGER NOT NULL,
    iteration_count INTEGER NOT NULL,
    symmetry_level INTEGER NOT NULL,
    gamma DOUBLE PRECISION NOT NULL,
    image TEXT
);

CREATE TABLE functions
(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    fractal_id BIGINT NOT NULL,
    weight DOUBLE PRECISION NOT NULL,

    FOREIGN KEY (fractal_id) REFERENCES fractals (id) ON DELETE CASCADE
);

CREATE TABLE affine_params
(
    id BIGSERIAL PRIMARY KEY,
    fractal_id BIGSERIAL NOT NULL,
    color VARCHAR(7) NOT NULL,
    a DOUBLE PRECISION NOT NULL,
    b DOUBLE PRECISION NOT NULL,
    c DOUBLE PRECISION NOT NULL,
    d DOUBLE PRECISION NOT NULL,
    e DOUBLE PRECISION NOT NULL,
    f DOUBLE PRECISION NOT NULL,

    FOREIGN KEY (fractal_id) REFERENCES fractals (id) ON DELETE CASCADE
);
