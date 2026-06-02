package api

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"

	mockery "github.com/mizentui/fractal-flame/iam/internal/api/v1/mock"
	authctx "github.com/mizentui/fractal-flame/iam/internal/auth"
	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
	iamv1 "github.com/mizentui/fractal-flame/iam/pkg/proto/v1"
)

const (
	userID           int64 = 1
	emptyUserID      int64 = 0
	username               = "mizentui"
	updatedUsername        = "mizentui-new"
	password               = "1238124AAs"
	tooLongPassword        = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	accessToken            = "some_access_token"
	refreshToken           = "some_refresh_token"
	imageName              = "avatar.png"
	image                  = "some image"
	imageBytesString       = "some image bytes"
)

var (
	ErrAuthService  = errors.New("some auth service error")
	ErrUserService  = errors.New("some user service error")
	ErrImageService = errors.New("some image service error")
	imageBytes      = []byte(imageBytesString)
)

func TestRegister(t *testing.T) {
	type args struct {
		ctx context.Context
		req *iamv1.AuthRequest
	}

	tests := []struct {
		message string
		args    args
		want    *iamv1.RegisterResponse
		err     error
		mock    func(*mockery.AuthServiceMock, *mockery.UserServiceMock, *mockery.ImageServiceMock, args)
	}{
		{
			message: "validation error: empty username",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: "", Password: password},
			},
			want: nil,
			err:  errs.ErrInvalidCredentials,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.AssertNotCalled(t, "Register", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "validation error: too long password",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: username, Password: tooLongPassword},
			},
			want: nil,
			err:  errs.ErrInvalidCredentials,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.AssertNotCalled(t, "Register", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "auth service error: failed to register",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: username, Password: password},
			},
			want: nil,
			err:  ErrAuthService,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.On("Register", a.ctx, a.req.Username, a.req.Password).Once().Return(emptyUserID, ErrAuthService)
			},
		},
		{
			message: "auth service ok: successfully register",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: username, Password: password},
			},
			want: &iamv1.RegisterResponse{UserId: userID},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.On("Register", a.ctx, a.req.Username, a.req.Password).Once().Return(userID, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			authService := mockery.NewAuthServiceMock(t)
			userService := mockery.NewUserServiceMock(t)
			imageService := mockery.NewImageServiceMock(t)

			test.mock(authService, userService, imageService, test.args)

			api := New(authService, userService, imageService)

			resp, err := api.Register(test.args.ctx, test.args.req)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, resp)
		})
	}
}

func TestLogin(t *testing.T) {
	type args struct {
		ctx context.Context
		req *iamv1.AuthRequest
	}

	tests := []struct {
		message string
		args    args
		want    *iamv1.LoginResponse
		err     error
		mock    func(*mockery.AuthServiceMock, *mockery.UserServiceMock, *mockery.ImageServiceMock, args)
	}{
		{
			message: "validation error: empty username",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: "", Password: password},
			},
			want: nil,
			err:  errs.ErrInvalidCredentials,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.AssertNotCalled(t, "Login", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "validation error: too long password",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: username, Password: tooLongPassword},
			},
			want: nil,
			err:  errs.ErrInvalidCredentials,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.AssertNotCalled(t, "Login", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "auth service error: failed to login",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: username, Password: password},
			},
			want: nil,
			err:  ErrAuthService,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.On("Login", a.ctx, a.req.Username, a.req.Password).Once().Return(model.TokenPair{}, ErrAuthService)
			},
		},
		{
			message: "auth service ok: successfully login",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: username, Password: password},
			},
			want: &iamv1.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				pair := model.TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}

				asm.On("Login", a.ctx, a.req.Username, a.req.Password).Once().Return(pair, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			authService := mockery.NewAuthServiceMock(t)
			userService := mockery.NewUserServiceMock(t)
			imageService := mockery.NewImageServiceMock(t)

			test.mock(authService, userService, imageService, test.args)

			api := New(authService, userService, imageService)

			resp, err := api.Login(test.args.ctx, test.args.req)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, resp)
		})
	}
}

func TestRefresh(t *testing.T) {
	type args struct {
		ctx context.Context
		req *iamv1.RefreshRequest
	}

	tests := []struct {
		message string
		args    args
		want    *iamv1.RefreshResponse
		err     error
		mock    func(*mockery.AuthServiceMock, *mockery.UserServiceMock, *mockery.ImageServiceMock, args)
	}{
		{
			message: "validation error: too long refresh token",
			args: args{
				ctx: context.Background(),
				req: &iamv1.RefreshRequest{RefreshToken: tooLongPassword},
			},
			want: nil,
			err:  errs.ErrInvalidToken,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.AssertNotCalled(t, "Refresh", mock.Anything, mock.Anything)
			},
		},
		{
			message: "auth service error: failed to refresh",
			args: args{
				ctx: context.Background(),
				req: &iamv1.RefreshRequest{RefreshToken: refreshToken},
			},
			want: nil,
			err:  ErrAuthService,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.On("Refresh", a.ctx, a.req.RefreshToken).Once().Return(model.TokenPair{}, ErrAuthService)
			},
		},
		{
			message: "auth service ok: successfully refresh",
			args: args{
				ctx: context.Background(),
				req: &iamv1.RefreshRequest{RefreshToken: refreshToken},
			},
			want: &iamv1.RefreshResponse{AccessToken: accessToken, RefreshToken: refreshToken},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				pair := model.TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}

				asm.On("Refresh", a.ctx, a.req.RefreshToken).Once().Return(pair, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			authService := mockery.NewAuthServiceMock(t)
			userService := mockery.NewUserServiceMock(t)
			imageService := mockery.NewImageServiceMock(t)

			test.mock(authService, userService, imageService, test.args)

			api := New(authService, userService, imageService)

			resp, err := api.Refresh(test.args.ctx, test.args.req)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, resp)
		})
	}
}

func TestGetUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *iamv1.GetUserRequest
	}

	tests := []struct {
		message string
		args    args
		want    *iamv1.GetUserResponse
		err     error
		mock    func(*mockery.AuthServiceMock, *mockery.UserServiceMock, *mockery.ImageServiceMock, args)
	}{
		{
			message: "validation error: empty user id",
			args: args{
				ctx: context.Background(),
				req: &iamv1.GetUserRequest{UserId: emptyUserID},
			},
			want: nil,
			err:  errs.ErrInvalidUserID,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				usm.AssertNotCalled(t, "GetUser", mock.Anything, mock.Anything)
			},
		},
		{
			message: "user service error: failed to get user",
			args: args{
				ctx: context.Background(),
				req: &iamv1.GetUserRequest{UserId: userID},
			},
			want: nil,
			err:  ErrUserService,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				usm.On("GetUser", a.ctx, a.req.UserId).Once().Return(model.User{}, ErrUserService)
			},
		},
		{
			message: "user service ok: successfully get user",
			args: args{
				ctx: context.Background(),
				req: &iamv1.GetUserRequest{UserId: userID},
			},
			want: &iamv1.GetUserResponse{User: &iamv1.User{UserId: userID, Username: username, Image: imageName}},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				user := model.User{ID: userID, Username: username, Image: imageName}

				usm.On("GetUser", a.ctx, a.req.UserId).Once().Return(user, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			authService := mockery.NewAuthServiceMock(t)
			userService := mockery.NewUserServiceMock(t)
			imageService := mockery.NewImageServiceMock(t)

			test.mock(authService, userService, imageService, test.args)

			api := New(authService, userService, imageService)

			resp, err := api.GetUser(test.args.ctx, test.args.req)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, resp)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *iamv1.UpdateUserRequest
	}

	tests := []struct {
		message string
		args    args
		want    *iamv1.UpdateUserResponse
		err     error
		mock    func(*mockery.AuthServiceMock, *mockery.UserServiceMock, *mockery.ImageServiceMock, args)
	}{
		{
			message: "user service error: nothing to update",
			args: args{
				ctx: authctx.WithUserID(context.Background(), userID),
				req: &iamv1.UpdateUserRequest{},
			},
			want: nil,
			err:  errs.ErrNothingToUpdate,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				usm.On("UpdateUser", a.ctx, userID, "", "", "").Once().Return(model.User{}, errs.ErrNothingToUpdate)
			},
		},
		{
			message: "context error: failed to extract user id",
			args: args{
				ctx: context.Background(),
				req: &iamv1.UpdateUserRequest{Username: wrapperspb.String(updatedUsername)},
			},
			want: nil,
			err:  errs.ErrInvalidCtxValue,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				usm.AssertNotCalled(t, "UpdateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "user service error: failed to update user",
			args: args{
				ctx: authctx.WithUserID(context.Background(), userID),
				req: &iamv1.UpdateUserRequest{
					Username: wrapperspb.String(updatedUsername),
					Password: wrapperspb.String(password),
					Image:    wrapperspb.String(image),
				},
			},
			want: nil,
			err:  ErrUserService,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				usm.On("UpdateUser", a.ctx, userID, updatedUsername, password, image).Once().Return(model.User{}, ErrUserService)
			},
		},
		{
			message: "user service ok: successfully update user",
			args: args{
				ctx: authctx.WithUserID(context.Background(), userID),
				req: &iamv1.UpdateUserRequest{
					Username: wrapperspb.String(updatedUsername),
					Password: wrapperspb.String(password),
					Image:    wrapperspb.String(image),
				},
			},
			want: &iamv1.UpdateUserResponse{User: &iamv1.User{UserId: userID, Username: updatedUsername, Image: imageName}},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				user := model.User{ID: userID, Username: updatedUsername, Image: imageName}

				usm.On("UpdateUser", a.ctx, userID, updatedUsername, password, image).Once().Return(user, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			authService := mockery.NewAuthServiceMock(t)
			userService := mockery.NewUserServiceMock(t)
			imageService := mockery.NewImageServiceMock(t)

			test.mock(authService, userService, imageService, test.args)

			api := New(authService, userService, imageService)

			resp, err := api.UpdateUser(test.args.ctx, test.args.req)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, resp)
		})
	}
}

func TestGetImage(t *testing.T) {
	type args struct {
		ctx context.Context
		req *iamv1.GetImageRequest
	}

	tests := []struct {
		message string
		args    args
		want    *iamv1.GetImageResponse
		err     error
		mock    func(*mockery.AuthServiceMock, *mockery.UserServiceMock, *mockery.ImageServiceMock, args)
	}{
		{
			message: "validation error: empty image name",
			args: args{
				ctx: context.Background(),
				req: &iamv1.GetImageRequest{Name: ""},
			},
			want: nil,
			err:  errs.ErrInvalidImageName,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				ism.AssertNotCalled(t, "GetImage", mock.Anything, mock.Anything)
			},
		},
		{
			message: "image service error: failed to get image",
			args: args{
				ctx: context.Background(),
				req: &iamv1.GetImageRequest{Name: imageName},
			},
			want: nil,
			err:  ErrImageService,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				ism.On("GetImage", a.ctx, a.req.Name).Once().Return(nil, ErrImageService)
			},
		},
		{
			message: "image service ok: successfully get image",
			args: args{
				ctx: context.Background(),
				req: &iamv1.GetImageRequest{Name: imageName},
			},
			want: &iamv1.GetImageResponse{Image: imageBytes},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				ism.On("GetImage", a.ctx, a.req.Name).Once().Return(imageBytes, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			authService := mockery.NewAuthServiceMock(t)
			userService := mockery.NewUserServiceMock(t)
			imageService := mockery.NewImageServiceMock(t)

			test.mock(authService, userService, imageService, test.args)

			api := New(authService, userService, imageService)

			resp, err := api.GetImage(test.args.ctx, test.args.req)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, resp)
		})
	}
}
