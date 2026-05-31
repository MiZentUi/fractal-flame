package image

import (
	"context"
	"fmt"
)

type ImageRepository interface {
	FindByName(ctx context.Context, name string) ([]byte, error)
}

type service struct {
	repository ImageRepository
}

func New(repository ImageRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetImage(ctx context.Context, name string) ([]byte, error) {
	image, err := s.repository.FindByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("find image by name: %w", err)
	}

	return image, nil
}
