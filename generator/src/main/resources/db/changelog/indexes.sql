--liquibase formatted sql

--changeset mizentui:indexes

CREATE INDEX fractals_user_id_idx ON fractals (user_id);
CREATE INDEX functions_fractal_id_idx ON functions (fractal_id);
CREATE INDEX affine_params_fractal_id_idx ON affine_params (fractal_id);
