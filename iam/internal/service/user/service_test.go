package user

import (
	"context"
	"encoding/base64"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
	mockery "github.com/mizentui/fractal-flame/iam/internal/service/user/mock"
)

const (
	invalidBase64Image = "some image"
)

var (
	validImageBytes     = []byte("some image bytes")
	validBase64Image    = base64.StdEncoding.EncodeToString(validImageBytes)
	errUserRepository   = errors.New("some user repo error")
	errImageRepository  = errors.New("some image repo error")
	errPasswordHasher   = errors.New("some hasher error")
	errImageValidator   = errors.New("some image validator error")
	errPasswordValidate = errors.New("password is too short")
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
		mock    func(*mockery.UserRepositoryMock, *mockery.ImageRepositoryMock, *mockery.PasswordValidatorMock, *mockery.PasswordHasherMock, *mockery.ImageValidatorMock, args)
	}{
		{
			message: "user repository error: failed to find user",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: model.User{},
			err:  errUserRepository,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				urm.On("FindByID", a.ctx, a.id).Once().Return(model.User{}, errUserRepository)
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
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
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
			pwdValidator := mockery.NewPasswordValidatorMock(t)
			hasher := mockery.NewPasswordHasherMock(t)
			imgValidator := mockery.NewImageValidatorMock(t)

			test.mock(userRepo, imageRepo, pwdValidator, hasher, imgValidator, test.args)

			service := New(userRepo, imageRepo, pwdValidator, hasher, imgValidator)

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
		message     string
		args        args
		want        model.User
		err         error
		errContains string
		mock        func(*mockery.UserRepositoryMock, *mockery.ImageRepositoryMock, *mockery.PasswordValidatorMock, *mockery.PasswordHasherMock, *mockery.ImageValidatorMock, args)
	}{
		{
			message: "validation error: nothing to update",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: model.User{},
			err:  errs.ErrNothingToUpdate,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				pvm.AssertNotCalled(t, "Validate", mock.Anything)
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				ivm.AssertNotCalled(t, "Validate", mock.Anything)
				irm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything)
				urm.AssertNotCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "password validator error: weak password",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui",
				password: "123",
				image:    validBase64Image,
			},
			want: model.User{},
			err:  errs.ErrWeakPassword,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				pvm.On("Validate", a.password).Once().Return(errPasswordValidate)
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				ivm.AssertNotCalled(t, "Validate", mock.Anything)
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
				image:    validBase64Image,
			},
			want: model.User{},
			err:  errPasswordHasher,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return("", errPasswordHasher)
				ivm.AssertNotCalled(t, "Validate", mock.Anything)
				irm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything)
				urm.AssertNotCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "image decode error: invalid base64 image",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui",
				image:    invalidBase64Image,
			},
			want: model.User{},
			err:  errs.ErrInvalidImage,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				pvm.AssertNotCalled(t, "Validate", mock.Anything)
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				ivm.AssertNotCalled(t, "Validate", mock.Anything)
				irm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything)
				urm.AssertNotCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "image validator error: invalid image bytes",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui",
				image:    validBase64Image,
			},
			want:        model.User{},
			errContains: "validate image",
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				pvm.AssertNotCalled(t, "Validate", mock.Anything)
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				ivm.On("Validate", validImageBytes).Once().Return(errImageValidator)
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
				image:    validBase64Image,
			},
			want: model.User{},
			err:  errImageRepository,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				hash := "some_password_hash"

				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(hash, nil)
				ivm.On("Validate", validImageBytes).Once().Return(nil)
				irm.On("Save", a.ctx, validImageBytes).Once().Return("", errImageRepository)
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
				image:    validBase64Image,
			},
			want: model.User{},
			err:  errUserRepository,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				hash, imageName := "some_password_hash", "avatar.png"

				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(hash, nil)
				ivm.On("Validate", validImageBytes).Once().Return(nil)
				irm.On("Save", a.ctx, validImageBytes).Once().Return(imageName, nil)
				urm.On("Update", a.ctx, a.id, a.username, hash, imageName).Once().Return(model.User{}, errUserRepository)
			},
		},
		{
			message: "user repository ok: update only username",
			args: args{
				ctx:      context.Background(),
				id:       1,
				username: "mizentui-new",
			},
			want: model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: "avatar.png"},
			err:  nil,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				user := model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: "avatar.png"}

				pvm.AssertNotCalled(t, "Validate", mock.Anything)
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				ivm.AssertNotCalled(t, "Validate", mock.Anything)
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
			},
			want: model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: "avatar.png"},
			err:  nil,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				hash := "some_password_hash"
				user := model.User{ID: 1, Username: "mizentui-new", Password: hash, Image: "avatar.png"}

				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(hash, nil)
				ivm.AssertNotCalled(t, "Validate", mock.Anything)
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
				image:    validBase64Image,
			},
			want: model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: "avatar.png"},
			err:  nil,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				imageName := "avatar.png"
				user := model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: imageName}

				pvm.AssertNotCalled(t, "Validate", mock.Anything)
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				ivm.On("Validate", validImageBytes).Once().Return(nil)
				irm.On("Save", a.ctx, validImageBytes).Once().Return(imageName, nil)
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
				image:    validBase64Image,
			},
			want: model.User{ID: 1, Username: "mizentui-new", Password: "some_password_hash", Image: "avatar.png"},
			err:  nil,
			mock: func(urm *mockery.UserRepositoryMock, irm *mockery.ImageRepositoryMock, pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, ivm *mockery.ImageValidatorMock, a args) {
				hash, imageName := "some_password_hash", "avatar.png"
				user := model.User{ID: 1, Username: "mizentui-new", Password: hash, Image: imageName}

				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(hash, nil)
				ivm.On("Validate", validImageBytes).Once().Return(nil)
				irm.On("Save", a.ctx, validImageBytes).Once().Return(imageName, nil)
				urm.On("Update", a.ctx, a.id, a.username, hash, imageName).Once().Return(user, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			userRepo := mockery.NewUserRepositoryMock(t)
			imageRepo := mockery.NewImageRepositoryMock(t)
			pwdValidator := mockery.NewPasswordValidatorMock(t)
			hasher := mockery.NewPasswordHasherMock(t)
			imgValidator := mockery.NewImageValidatorMock(t)

			test.mock(userRepo, imageRepo, pwdValidator, hasher, imgValidator, test.args)

			service := New(userRepo, imageRepo, pwdValidator, hasher, imgValidator)

			user, err := service.UpdateUser(
				test.args.ctx,
				test.args.id,
				test.args.username,
				test.args.password,
				test.args.image,
			)
			if err != nil {
				require.Error(t, err)
				if test.err != nil {
					require.ErrorIs(t, err, test.err)
				}
				if test.errContains != "" {
					require.ErrorContains(t, err, test.errContains)
				}

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, user)
		})
	}
}
