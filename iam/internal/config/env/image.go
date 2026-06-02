package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type imageEnvConfig struct {
	Width  int `env:"IMAGE_WIDTH,required"`
	Height int `env:"IMAGE_HEIGHT,required"`
}

type imageConfig struct {
	raw imageEnvConfig
}

func (ic *imageConfig) Width() int {
	return ic.raw.Width
}

func (ic *imageConfig) Height() int {
	return ic.raw.Height
}

func NewImageConfig() (*imageConfig, error) {
	var raw imageEnvConfig

	err := env.Parse(&raw)
	if err != nil {
		return nil, fmt.Errorf("load image envs: %w", err)
	}

	return &imageConfig{
		raw: raw,
	}, nil
}
