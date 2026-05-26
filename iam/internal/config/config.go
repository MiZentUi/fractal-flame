package config

import (
	"fmt"

	"github.com/mizentui/fractal-flame/iam/internal/config/env"
)

type Logger interface {
	Level() string
	AsJSON() bool
}

type GRPC interface {
	Address() string
}

type Postgres interface {
	URI() string
}

type Minio interface {
	Endpoint() string
	AccessKey() string
	SecretKey() string
}

type Auth interface {
	PasswordEntropy() float64
	BcryptCost() int
	SigningKey() string
}

type config struct {
	GRPC     GRPC
	Logger   Logger
	Postgres Postgres
	Minio    Minio
	Auth     Auth
}

var app *config

func App() *config {
	return app
}

func Setup() error {
	grpc, err := env.NewGRPCConfig()
	if err != nil {
		return fmt.Errorf("setup grpc config: %w", err)
	}

	logger, err := env.NewLoggerConfig()
	if err != nil {
		return fmt.Errorf("setup grpc config: %w", err)
	}

	postgres, err := env.NewPostgresConfig()
	if err != nil {
		return fmt.Errorf("setup grpc config: %w", err)
	}

	minio, err := env.NewMinioConfig()
	if err != nil {
		return fmt.Errorf("setup grpc config: %w", err)
	}

	auth, err := env.NewAuthConfig()
	if err != nil {
		return fmt.Errorf("setup auth config: %w", err)
	}

	app = &config{
		GRPC:     grpc,
		Logger:   logger,
		Postgres: postgres,
		Minio:    minio,
		Auth:     auth,
	}

	return nil
}
