package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type authEnvConfig struct {
	PasswordEntropy float64 `env:"MIN_PASSWORD_ENTROPY,required"`
	BcryptCost      int     `env:"BCRYPT_COST,required"`
	SigningKey      string  `env:"JWT_SIGNING_KEY,required"`
}

type authConfig struct {
	raw authEnvConfig
}

func (ac *authConfig) PasswordEntropy() float64 {
	return ac.raw.PasswordEntropy
}

func (ac *authConfig) BcryptCost() int {
	return ac.raw.BcryptCost
}

func (ac *authConfig) SigningKey() string {
	return ac.raw.SigningKey
}

func NewAuthConfig() (*authConfig, error) {
	var raw authEnvConfig

	err := env.Parse(&raw)
	if err != nil {
		return nil, fmt.Errorf("load jwt envs: %w", err)
	}

	return &authConfig{
		raw: raw,
	}, nil
}
