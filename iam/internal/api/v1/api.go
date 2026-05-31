package api

import (
	"context"

	"github.com/mizentui/fractal-flame/iam/internal/api/v1/dto"
	"github.com/mizentui/fractal-flame/iam/internal/auth"
	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
	iamv1 "github.com/mizentui/fractal-flame/iam/pkg/proto/v1"
)

type AuthService interface {
	Register(ctx context.Context, username, password string) (int64, error)
	Login(ctx context.Context, username, password string) (model.TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (model.TokenPair, error)
}

type UserService interface {
	GetUser(ctx context.Context, id int64) (model.User, error)
	UpdateUser(ctx context.Context, id int64, username, password, image string) (model.User, error)
}

type ImageService interface {
	GetImage(ctx context.Context, name string) ([]byte, error)
}

var _ iamv1.IAMServiceServer = (*api)(nil)

type api struct {
	auth  AuthService
	user  UserService
	image ImageService
}

func New(auth AuthService, user UserService, image ImageService) *api {
	return &api{
		auth:  auth,
		user:  user,
		image: image,
	}
}

func (a *api) Register(ctx context.Context, req *iamv1.AuthRequest) (*iamv1.RegisterResponse, error) {
	if req.ValidateAll() != nil {
		return nil, errs.ErrInvalidCredentials
	}

	userID, err := a.auth.Register(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &iamv1.RegisterResponse{UserId: userID}, nil
}

func (a *api) Login(ctx context.Context, req *iamv1.AuthRequest) (*iamv1.LoginResponse, error) {
	if req.ValidateAll() != nil {
		return nil, errs.ErrInvalidCredentials
	}

	pair, err := a.auth.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &iamv1.LoginResponse{AccessToken: pair.AccessToken, RefreshToken: pair.RefreshToken}, nil
}

func (a *api) Refresh(ctx context.Context, req *iamv1.RefreshRequest) (*iamv1.RefreshResponse, error) {
	if req.ValidateAll() != nil {
		return nil, errs.ErrInvalidToken
	}

	pair, err := a.auth.Refresh(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &iamv1.RefreshResponse{AccessToken: pair.AccessToken, RefreshToken: pair.RefreshToken}, nil
}

func (a *api) GetUser(ctx context.Context, req *iamv1.GetUserRequest) (*iamv1.GetUserResponse, error) {
	if req.ValidateAll() != nil {
		return nil, errs.ErrInvalidUserID
	}

	user, err := a.user.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &iamv1.GetUserResponse{User: dto.ModelToProto(user)}, nil
}

func (a *api) UpdateUser(ctx context.Context, req *iamv1.UpdateUserRequest) (*iamv1.UpdateUserResponse, error) {
	id, err := auth.ExtractUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := a.user.UpdateUser(ctx, id, req.Username.GetValue(), req.Password.GetValue(), req.Image.GetValue())
	if err != nil {
		return nil, err
	}

	return &iamv1.UpdateUserResponse{User: dto.ModelToProto(user)}, nil
}

func (a *api) GetImage(ctx context.Context, req *iamv1.GetImageRequest) (*iamv1.GetImageResponse, error) {
	if req.Name == "" {
		return nil, errs.ErrInvalidImageName
	}

	image, err := a.image.GetImage(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &iamv1.GetImageResponse{Image: image}, nil
}
