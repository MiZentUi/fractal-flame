package user

import (
	"context"
	"errors"
	"testing"

	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
	mockery "github.com/mizentui/fractal-flame/iam/internal/service/user/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	ErrUserRepository  = errors.New("some user repo error")
	ErrImageRepository = errors.New("some image repo error")
	ErrPasswordHasher  = errors.New("some hasher error")
)

func TestGetUser(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		message string
		args    args
		want    model.User
		err     error
		mock    func(*mockery.UserRepositoryMock, *mockery.ImageRepositoryMock, *mockery.PasswordValidatorMock, *mockery.PasswordHasherMock, args)
	}{
		{
			message: "user repository error: failed to find user",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: model.User{},
			err:  ErrUserRepository,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, a args) {
				urm.On("FindByID", a.ctx, a.id).Once().Return(model.User{}, ErrUserRepository)
			},
		},
		{
			message: "user repository ok: successfully find user",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: model.User{ID: 1, Username: "mizentui", Password: "some_password_hash", Image: "avatar.png"},
			err:  nil,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, a args) {
				user := model.User{ID: 1, Username: "mizentui", Password: "some_password_hash", Image: "avatar.png"}

				urm.On("FindByID", a.ctx, a.id).Once().Return(user, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			userRepo := mockery.NewUserRepositoryMock(t)
			imageRepo := mockery.NewImageRepositoryMock(t)
			validator := mockery.NewPasswordValidatorMock(t)
			hasher := mockery.NewPasswordHasherMock(t)

			test.mock(userRepo, imageRepo, validator, hasher, test.args)

			service := New(userRepo, imageRepo, validator, hasher)

			user, err := service.GetUser(test.args.ctx, test.args.id)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, user)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		ctx      context.Context
		id       int64
		username string
		password string
		image    string
	}

	tests := []struct {
		message string
		args    args
		want    model.User
		err     error
		mock    func(*mockery.UserRepositoryMock, *mockery.ImageRepositoryMock, *mockery.PasswordValidatorMock, *mockery.PasswordHasherMock, args)
	}{
		{
			message: "validator error: weak password",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui",
				password: "123",
				image:    "some image",
			},
			want: model.User{},
			err:  errs.ErrWeakPassword,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, a args) {
				pvm.On("Validate", a.password).Once().Return(errors.New("password is too short"))
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				irm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything)
				urm.AssertNotCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "password hasher error: failed to get hash",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui",
				password: "1238124AAs",
				image:    "some image",
			},
			want: model.User{},
			err:  ErrPasswordHasher,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, a args) {
				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return("", ErrPasswordHasher)
				irm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything)
				urm.AssertNotCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "image repository error: failed to save image",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui",
				password: "1238124AAs",
				image:    "some image",
			},
			want: model.User{},
			err:  ErrImageRepository,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, a args) {
				hash := "some_password_hash"

				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(hash, nil)
				irm.On("Save", a.ctx, a.image).Once().Return("", ErrImageRepository)
				urm.AssertNotCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "user repository error: failed to update user",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui",
				password: "1238124AAs",
				image:    "some image",
			},
			want: model.User{},
			err:  ErrUserRepository,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, a args) {
				hash, imageName := "some_password_hash", "avatar.png"

				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(hash, nil)
				irm.On("Save", a.ctx, a.image).Once().Return(imageName, nil)
				urm.On("Update", a.ctx, a.id, a.username, hash, imageName).Once().Return(model.User{}, ErrUserRepository)
			},
		},
		{
			message: "user repository ok: update only username",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui-new",
				password: "",
				image:    "",
			},
			want: model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: "avatar.png"},
			err:  nil,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, a args) {
				user := model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: "avatar.png"}

				pvm.AssertNotCalled(t, "Validate", mock.Anything)
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				irm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything)
				urm.On("Update", a.ctx, a.id, a.username, "", "").Once().Return(user, nil)
			},
		},
		{
			message: "user repository ok: update username and password",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui-new",
				password: "1238124AAs",
				image:    "",
			},
			want: model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: "avatar.png"},
			err:  nil,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, a args) {
				hash := "some_password_hash"
				user := model.User{ID: 1, Username: "mizentui-new", Password: hash, Image: "avatar.png"}

				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(hash, nil)
				irm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything)
				urm.On("Update", a.ctx, a.id, a.username, hash, "").Once().Return(user, nil)
			},
		},
		{
			message: "user repository ok: update username and image",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui-new",
				password: "",
				image:    "some image",
			},
			want: model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: "avatar.png"},
			err:  nil,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, a args) {
				imageName := "avatar.png"
				user := model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: imageName}

				pvm.AssertNotCalled(t, "Validate", mock.Anything)
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				irm.On("Save", a.ctx, a.image).Once().Return(imageName, nil)
				urm.On("Update", a.ctx, a.id, a.username, "", imageName).Once().Return(user, nil)
			},
		},
		{
			message: "user repository ok: update username, password and image",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui-new",
				password: "1238124AAs",
				image:    "some image",
			},
			want: model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: "avatar.png"},
			err:  nil,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, a args) {
				hash, imageName := "some_password_hash", "avatar.png"
				user := model.User{ID: 1, Username: "mizentui-new", Password: hash, Image: imageName}

				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(hash, nil)
				irm.On("Save", a.ctx, a.image).Once().Return(imageName, nil)
				urm.On("Update", a.ctx, a.id, a.username, hash, imageName).Once().Return(user, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			userRepo := mockery.NewUserRepositoryMock(t)
			imageRepo := mockery.NewImageRepositoryMock(t)
			validator := mockery.NewPasswordValidatorMock(t)
			hasher := mockery.NewPasswordHasherMock(t)

			test.mock(userRepo, imageRepo, validator, hasher, test.args)

			service := New(userRepo, imageRepo, validator, hasher)

			user, err := service.UpdateUser(
				test.args.ctx,
				test.args.id,
				test.args.username,
				test.args.password,
				test.args.image,
			)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, user)
		})
	}
}
