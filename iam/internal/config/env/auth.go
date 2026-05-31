package env

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type authEnvConfig struct {
	PasswordEntropy   float64       `env:"MIN_PASSWORD_ENTROPY,required"`
	BcryptCost        int           `env:"BCRYPT_COST,required"`
	AccessSigningKey  string        `env:"ACCESS_SIGNING_KEY,required"`
	RefreshSigningKey string        `env:"REFRESH_SIGNING_KEY,required"`
	AccessTokenTTL    time.Duration `env:"ACCESS_TOKEN_TTL,required"`
	RefreshTokenTTL   time.Duration `env:"REFRESH_TOKEN_TTL,required"`
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

func (ac *authConfig) AccessSigningKey() string {
	return ac.raw.AccessSigningKey
}

func (ac *authConfig) RefreshSigningKey() string {
	return ac.raw.RefreshSigningKey
}

func (ac *authConfig) AccessTokenTTL() time.Duration {
	return ac.raw.AccessTokenTTL
}

func (ac *authConfig) RefreshTokenTTL() time.Duration {
	return ac.raw.RefreshTokenTTL
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
