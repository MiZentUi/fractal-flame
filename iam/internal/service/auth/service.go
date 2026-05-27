package auth

import (
	"context"
	"errors"
	"fmt"

	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
)

type UserRepository interface {
	Save(ctx context.Context, username, password string) (int64, error)
	FindByUsername(ctx context.Context, username string) (model.User, error)
	FindByID(ctx context.Context, id int64) (model.User, error)
}

type PasswordValidator interface {
	Validate(password string) error
}

type PasswordHasher interface {
	HashAndSalt(password string) (string, error)
	ComparePasswords(hash, password string) error
}

type TokenManager interface {
	GenerateTokensPair(user model.User) (model.TokenPair, error)
	ValidateRefreshToken(token string) (int64, error)
}

type service struct {
	repository UserRepository
	validator  PasswordValidator
	hasher     PasswordHasher
	manager    TokenManager
}

func New(repository UserRepository, validator PasswordValidator, hasher PasswordHasher, manager TokenManager) *service {
	return &service{
		repository: repository,
		validator:  validator,
		hasher:     hasher,
		manager:    manager,
	}
}

func (s *service) Register(ctx context.Context, username, password string) (int64, error) {
	err := s.validator.Validate(password)
	if err != nil {
		return 0, errors.Join(errs.ErrWeakPassword, err)
	}

	hash, err := s.hasher.HashAndSalt(password)
	if err != nil {
		return 0, fmt.Errorf("get password hash: %w", err)
	}

	userID, err := s.repository.Save(ctx, username, hash)
	if err != nil {
		return 0, fmt.Errorf("add user to db: %w", err)
	}

	return userID, nil
}

func (s *service) Login(ctx context.Context, username, password string) (model.TokenPair, error) {
	user, err := s.repository.FindByUsername(ctx, username)
	if err != nil {
		return model.TokenPair{}, errs.ErrInvalidCredentials
	}

	err = s.hasher.ComparePasswords(user.Password, password)
	if err != nil {
		return model.TokenPair{}, errs.ErrInvalidCredentials
	}

	pair, err := s.manager.GenerateTokensPair(user)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("generate token pair: %w", err)
	}

	return pair, nil
}

func (s *service) Refresh(ctx context.Context, refreshToken string) (model.TokenPair, error) {
	id, err := s.manager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("validate refresh token: %w", err)
	}

	user, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("find user by id: %w", err)
	}

	pair, err := s.manager.GenerateTokensPair(user)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("generate access and refresh tokens: %w", err)
	}

	return pair, nil
}
