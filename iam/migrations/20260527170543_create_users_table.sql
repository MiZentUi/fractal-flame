-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL,
    image VARCHAR(50)
);

-- +goose Down
DROP TABLE IF EXISTS users;
