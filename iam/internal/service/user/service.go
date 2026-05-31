package user

import (
	"context"
	"encoding/base64"
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
	Save(ctx context.Context, bytes []byte) (string, error)
}

type PasswordValidator interface {
	Validate(password string) error
}

type PasswordHasher interface {
	HashAndSalt(password string) (string, error)
}

type ImageValidator interface {
	Validate(image []byte) error
}

type service struct {
	user         UserRepository
	image        ImageRepository
	pwdValidator PasswordValidator
	hasher       PasswordHasher
	imgValidator ImageValidator
}

func New(user UserRepository, image ImageRepository, pwdValidator PasswordValidator, hasher PasswordHasher, imgValidator ImageValidator) *service {
	return &service{
		user:         user,
		image:        image,
		pwdValidator: pwdValidator,
		hasher:       hasher,
		imgValidator: imgValidator,
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

	hash, err := s.getPasswordHash(password)
	if err != nil {
		return model.User{}, err
	}

	bytes, err := s.getImageBytes(image)
	if err != nil {
		return model.User{}, err
	}

	var name string
	if bytes != nil {
		name, err = s.image.Save(ctx, bytes)
		if err != nil {
			return model.User{}, fmt.Errorf("save avatar image: %w", err)
		}
	}

	user, err := s.user.Update(ctx, id, username, hash, name)
	if err != nil {
		return model.User{}, fmt.Errorf("update user: %w", err)
	}

	return user, nil
}

func (s *service) getPasswordHash(password string) (string, error) {
	if password == "" {
		return "", nil
	}

	err := s.pwdValidator.Validate(password)
	if err != nil {
		return "", errors.Join(errs.ErrWeakPassword, err)
	}

	hash, err := s.hasher.HashAndSalt(password)
	if err != nil {
		return "", fmt.Errorf("get password hash: %w", err)
	}

	return hash, nil
}

func (s *service) getImageBytes(image string) ([]byte, error) {
	if image == "" {
		return nil, nil
	}

	bytes, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		slog.Error("Failed to decode image", "err", err)

		return nil, errors.Join(errs.ErrInvalidImage, err)
	}

	err = s.imgValidator.Validate(bytes)
	if err != nil {
		return nil, fmt.Errorf("validate image", err)
	}

	return bytes, nil
}
