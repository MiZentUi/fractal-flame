package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type jwtEnvConfig struct {
	SigningKey string `env:"JWT_SIGNING_KEY,required"`
}

type jwtConfig struct {
	raw jwtEnvConfig
}

func (jc *jwtConfig) SigningKey() string {
	return jc.raw.SigningKey
}

func NewJWTConfig() (*jwtConfig, error) {
	var raw jwtEnvConfig

	err := env.Parse(&raw)
	if err != nil {
		return nil, fmt.Errorf("load jwt envs: %w", err)
	}

	return &jwtConfig{
		raw: raw,
	}, nil
}
