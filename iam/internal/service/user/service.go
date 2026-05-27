package user

import (
	"context"
	"errors"
	"fmt"

	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (model.User, error)
	Update(ctx context.Context, id int64, username, password, image string) (model.User, error)
}

type PasswordValidator interface {
	Validate(password string) error
}

type PasswordHasher interface {
	HashAndSalt(password string) (string, error)
}

type service struct {
	repository UserRepository
	validator  PasswordValidator
	hasher     PasswordHasher
}

func New(repository UserRepository, validator PasswordValidator, hasher PasswordHasher) *service {
	return &service{
		repository: repository,
		validator:  validator,
		hasher:     hasher,
	}
}

func (s *service) GetUser(ctx context.Context, id int64) (model.User, error) {
	user, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("find user by id: %w", err)
	}

	return user, nil
}

func (s *service) UpdateUser(ctx context.Context, id int64, username, password, image string) (model.User, error) {
	var hash string
	if password != "" {
		err := s.validator.Validate(password)
		if err != nil {
			return model.User{}, errors.Join(errs.ErrWeakPassword, err)
		}

		hash, err = s.hasher.HashAndSalt(password)
		if err != nil {
			return model.User{}, fmt.Errorf("get password hash: %w", err)
		}
	}

	user, err := s.repository.Update(ctx, id, username, hash, image)
	if err != nil {
		return model.User{}, fmt.Errorf("update user: %w", err)
	}

	return user, nil
}
