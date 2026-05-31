package env

import (
	"fmt"
	"net"

	"github.com/caarlos0/env/v11"
)

type minioEnvConfig struct {
	Host      string `env:"MINIO_HOST,required"`
	Port      string `env:"MINIO_PORT,required"`
	AccessKey string `env:"MINIO_ROOT_USER,required"`
	SecretKey string `env:"MINIO_ROOT_PASSWORD,required"`
}

type minioConfig struct {
	raw minioEnvConfig
}

func (mc *minioConfig) Endpoint() string {
	return net.JoinHostPort(mc.raw.Host, mc.raw.Port)
}

func (mc *minioConfig) AccessKey() string {
	return mc.raw.AccessKey
}

func (mc *minioConfig) SecretKey() string {
	return mc.raw.SecretKey
}

func NewMinioConfig() (*minioConfig, error) {
	var raw minioEnvConfig

	err := env.Parse(&raw)
	if err != nil {
		return nil, fmt.Errorf("load minio envs: %w", err)
	}

	return &minioConfig{
		raw: raw,
	}, nil
}
