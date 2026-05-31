package api

import (
	"context"
	"errors"
	"testing"

	mockery "github.com/mizentui/fractal-flame/iam/internal/api/v1/mock"
	authctx "github.com/mizentui/fractal-flame/iam/internal/auth"
	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
	iamv1 "github.com/mizentui/fractal-flame/iam/pkg/proto/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	ErrAuthService  = errors.New("some auth service error")
	ErrUserService  = errors.New("some user service error")
	ErrImageService = errors.New("some image service error")
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
				req: &iamv1.AuthRequest{Username: "", Password: "1238124AAs"},
			},
			want: nil,
			err:  errs.ErrInvalidCredentials,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.AssertNotCalled(t, "Register", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "validation error: empty password",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: "mizentui", Password: ""},
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
				req: &iamv1.AuthRequest{Username: "mizentui", Password: "1238124AAs"},
			},
			want: nil,
			err:  ErrAuthService,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.On("Register", a.ctx, a.req.Username, a.req.Password).Once().Return(int64(0), ErrAuthService)
			},
		},
		{
			message: "auth service ok: successfully register",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: "mizentui", Password: "1238124AAs"},
			},
			want: &iamv1.RegisterResponse{UserId: 1},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.On("Register", a.ctx, a.req.Username, a.req.Password).Once().Return(int64(1), nil)
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
				req: &iamv1.AuthRequest{Username: "", Password: "1238124AAs"},
			},
			want: nil,
			err:  errs.ErrInvalidCredentials,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				asm.AssertNotCalled(t, "Login", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "validation error: empty password",
			args: args{
				ctx: context.Background(),
				req: &iamv1.AuthRequest{Username: "mizentui", Password: ""},
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
				req: &iamv1.AuthRequest{Username: "mizentui", Password: "1238124AAs"},
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
				req: &iamv1.AuthRequest{Username: "mizentui", Password: "1238124AAs"},
			},
			want: &iamv1.LoginResponse{AccessToken: "some_access_token", RefreshToken: "some_refresh_token"},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				pair := model.TokenPair{AccessToken: "some_access_token", RefreshToken: "some_refresh_token"}

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
			message: "validation error: empty refresh token",
			args: args{
				ctx: context.Background(),
				req: &iamv1.RefreshRequest{RefreshToken: ""},
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
				req: &iamv1.RefreshRequest{RefreshToken: "some_refresh_token"},
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
				req: &iamv1.RefreshRequest{RefreshToken: "some_refresh_token"},
			},
			want: &iamv1.RefreshResponse{AccessToken: "some_access_token", RefreshToken: "some_refresh_token"},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				pair := model.TokenPair{AccessToken: "some_access_token", RefreshToken: "some_refresh_token"}

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
				req: &iamv1.GetUserRequest{UserId: 0},
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
				req: &iamv1.GetUserRequest{UserId: 1},
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
				req: &iamv1.GetUserRequest{UserId: 1},
			},
			want: &iamv1.GetUserResponse{User: &iamv1.User{UserId: 1, Username: "mizentui", Image: "avatar.png"}},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				user := model.User{ID: 1, Username: "mizentui", Image: "avatar.png"}

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
			message: "validation error: nothing to update",
			args: args{
				ctx: authctx.WithUserID(context.Background(), 1),
				req: &iamv1.UpdateUserRequest{},
			},
			want: nil,
			err:  errs.ErrNothingToUpdate,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				usm.AssertNotCalled(t, "UpdateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "context error: failed to extract user id",
			args: args{
				ctx: context.Background(),
				req: &iamv1.UpdateUserRequest{Username: wrapperspb.String("mizentui-new")},
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
				ctx: authctx.WithUserID(context.Background(), 1),
				req: &iamv1.UpdateUserRequest{
					Username: wrapperspb.String("mizentui-new"),
					Password: wrapperspb.String("1238124AAs"),
					Image:    wrapperspb.String("some image"),
				},
			},
			want: nil,
			err:  ErrUserService,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				usm.On("UpdateUser", a.ctx, int64(1), "mizentui-new", "1238124AAs", "some image").Once().Return(model.User{}, ErrUserService)
			},
		},
		{
			message: "user service ok: successfully update user",
			args: args{
				ctx: authctx.WithUserID(context.Background(), 1),
				req: &iamv1.UpdateUserRequest{
					Username: wrapperspb.String("mizentui-new"),
					Password: wrapperspb.String("1238124AAs"),
					Image:    wrapperspb.String("some image"),
				},
			},
			want: &iamv1.UpdateUserResponse{User: &iamv1.User{UserId: 1, Username: "mizentui-new", Image: "avatar.png"}},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				user := model.User{ID: 1, Username: "mizentui-new", Image: "avatar.png"}

				usm.On("UpdateUser", a.ctx, int64(1), "mizentui-new", "1238124AAs", "some image").Once().Return(user, nil)
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
				req: &iamv1.GetImageRequest{Name: "avatar.png"},
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
				req: &iamv1.GetImageRequest{Name: "avatar.png"},
			},
			want: &iamv1.GetImageResponse{Image: []byte("some image bytes")},
			err:  nil,
			mock: func(asm *mockery.AuthServiceMock, usm *mockery.UserServiceMock, ism *mockery.ImageServiceMock, a args) {
				image := []byte("some image bytes")

				ism.On("GetImage", a.ctx, a.req.Name).Once().Return(image, nil)
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
