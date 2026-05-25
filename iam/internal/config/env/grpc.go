package env

import (
	"fmt"
	"net"

	"github.com/caarlos0/env/v11"
)

type grpcEnvConfig struct {
	Host string `env:"IAM_HOST,required"`
	Port string `env:"IAM_PORT,required"`
}

type grpcConfig struct {
	raw grpcEnvConfig
}

func (gc *grpcConfig) Address() string {
	return net.JoinHostPort(gc.raw.Host, gc.raw.Port)
}

func NewGRPCConfig() (*grpcConfig, error) {
	var raw grpcEnvConfig

	err := env.Parse(&raw)
	if err != nil {
		return nil, fmt.Errorf("load gRPC envs: %w", err)
	}

	return &grpcConfig{
		raw: raw,
	}, nil
}
