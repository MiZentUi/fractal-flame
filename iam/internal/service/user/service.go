package user

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (model.User, error)
	Update(ctx context.Context, id int64, username, password, image string) (model.User, error)
}

type ImageRepository interface {
	Save(ctx context.Context, image string) (string, error)
}

type PasswordValidator interface {
	Validate(password string) error
}

type PasswordHasher interface {
	HashAndSalt(password string) (string, error)
}

type service struct {
	user      UserRepository
	image     ImageRepository
	validator PasswordValidator
	hasher    PasswordHasher
}

func New(user UserRepository, image ImageRepository, validator PasswordValidator, hasher PasswordHasher) *service {
	return &service{
		user:      user,
		image:     image,
		validator: validator,
		hasher:    hasher,
	}
}

func (s *service) GetUser(ctx context.Context, id int64) (model.User, error) {
	user, err := s.user.FindByID(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("find user by id: %w", err)
	}

	return user, nil
}

func (s *service) UpdateUser(ctx context.Context, id int64, username, password, image string) (model.User, error) {
	if username == "" && password == "" && image == "" {
		slog.Warn("Nothing user data to update", "err", errs.ErrNothingToUpdate)

		return model.User{}, errs.ErrNothingToUpdate
	}

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

	var imageName string
	if image != "" {
		name, err := s.image.Save(ctx, image)
		if err != nil {
			return model.User{}, fmt.Errorf("save avatar image: %w", err)
		}

		imageName = name
	}

	user, err := s.user.Update(ctx, id, username, hash, imageName)
	if err != nil {
		return model.User{}, fmt.Errorf("update user: %w", err)
	}

	return user, nil
}
