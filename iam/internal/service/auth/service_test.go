package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
	mockery "github.com/mizentui/fractal-flame/iam/internal/service/auth/mock"
)

var (
	ErrPasswordHasher = errors.New("some hasher error")
	ErrUserRepository = errors.New("some user repo error")
	ErrTokenManager   = errors.New("some token manager error")
)

func TestRegister(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		password string
	}

	tests := []struct {
		message string
		args    args
		want    int64
		err     error
		mock    func(*mockery.PasswordValidatorMock, *mockery.PasswordHasherMock, *mockery.UserRepositoryMock, *mockery.TokenManagerMock, args)
	}{
		{
			message: "validator error: short password",
			args: args{
				ctx:      context.Background(),
				username: "mizentui",
				password: "123",
			},
			want: 0,
			err:  errs.ErrWeakPassword,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				pvm.On("Validate", a.password).Once().Return(errors.New("password is too short"))
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				urm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "validator error: require special characters",
			args: args{
				ctx:      context.Background(),
				username: "mizentui",
				password: "123123123123123123",
			},
			want: 0,
			err:  errs.ErrWeakPassword,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				pvm.On("Validate", a.password).Once().Return(errors.New("require special characters in password"))
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				urm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "password hasher error: failed to get hash",
			args: args{
				ctx:      context.Background(),
				username: "mizentui",
				password: "1238124AAs",
			},
			want: 0,
			err:  ErrPasswordHasher,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return("", ErrPasswordHasher)
				urm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "user repository error: failed to save user",
			args: args{
				ctx:      context.Background(),
				username: "mizentui",
				password: "1238124AAs",
			},
			want: 0,
			err:  ErrUserRepository,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				hash, id := "some_password_hash", int64(0)

				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(hash, nil)
				urm.On("Save", a.ctx, a.username, hash).Once().Return(id, ErrUserRepository)
			},
		},
		{
			message: "user repository ok: successfully save user",
			args: args{
				ctx:      context.Background(),
				username: "mizentui",
				password: "1238124AAs",
			},
			want: 1,
			err:  nil,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				hash, id := "some_password_hash", int64(1)

				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(hash, nil)
				urm.On("Save", a.ctx, a.username, hash).Once().Return(id, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			repo := mockery.NewUserRepositoryMock(t)
			validator := mockery.NewPasswordValidatorMock(t)
			hasher := mockery.NewPasswordHasherMock(t)
			manager := mockery.NewTokenManagerMock(t)

			test.mock(validator, hasher, repo, manager, test.args)

			service := New(repo, validator, hasher, manager)

			id, err := service.Register(test.args.ctx, test.args.username, test.args.password)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.NotEmpty(t, id)
			require.Equal(t, id, test.want)
		})
	}
}

func TestLogin(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		password string
	}

	tests := []struct {
		message string
		args    args
		want    model.TokenPair
		err     error
		mock    func(*mockery.PasswordValidatorMock, *mockery.PasswordHasherMock, *mockery.UserRepositoryMock, *mockery.TokenManagerMock, args)
	}{
		{
			message: "user repository error: failed to find user",
			args: args{
				ctx:      context.Background(),
				username: "mizentui",
				password: "1238124AAs",
			},
			want: model.TokenPair{},
			err:  ErrUserRepository,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				urm.On("FindByUsername", a.ctx, a.username).Once().Return(model.User{}, ErrUserRepository)
				phm.AssertNotCalled(t, "ComparePasswords", mock.Anything, mock.Anything)
				tmm.AssertNotCalled(t, "GenerateTokensPair", mock.Anything)
			},
		},
		{
			message: "password hasher error: invalid credentials",
			args: args{
				ctx:      context.Background(),
				username: "mizentui",
				password: "1238124AAs",
			},
			want: model.TokenPair{},
			err:  errs.ErrInvalidCredentials,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				user := model.User{ID: 1, Username: a.username, Password: "some_password_hash"}

				urm.On("FindByUsername", a.ctx, a.username).Once().Return(user, nil)
				phm.On("ComparePasswords", user.Password, a.password).Once().Return(ErrPasswordHasher)
				tmm.AssertNotCalled(t, "GenerateTokensPair", mock.Anything)
			},
		},
		{
			message: "token manager error: failed to generate tokens",
			args: args{
				ctx:      context.Background(),
				username: "mizentui",
				password: "1238124AAs",
			},
			want: model.TokenPair{},
			err:  ErrTokenManager,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				user := model.User{ID: 1, Username: a.username, Password: "some_password_hash"}

				urm.On("FindByUsername", a.ctx, a.username).Once().Return(user, nil)
				phm.On("ComparePasswords", user.Password, a.password).Once().Return(nil)
				tmm.On("GenerateTokensPair", user).Once().Return(model.TokenPair{}, ErrTokenManager)
			},
		},
		{
			message: "token manager ok: successfully login",
			args: args{
				ctx:      context.Background(),
				username: "mizentui",
				password: "1238124AAs",
			},
			want: model.TokenPair{AccessToken: "some_access_token", RefreshToken: "some_refresh_token"},
			err:  nil,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				user := model.User{ID: 1, Username: a.username, Password: "some_password_hash"}
				pair := model.TokenPair{AccessToken: "some_access_token", RefreshToken: "some_refresh_token"}

				urm.On("FindByUsername", a.ctx, a.username).Once().Return(user, nil)
				phm.On("ComparePasswords", user.Password, a.password).Once().Return(nil)
				tmm.On("GenerateTokensPair", user).Once().Return(pair, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			repo := mockery.NewUserRepositoryMock(t)
			validator := mockery.NewPasswordValidatorMock(t)
			hasher := mockery.NewPasswordHasherMock(t)
			manager := mockery.NewTokenManagerMock(t)

			test.mock(validator, hasher, repo, manager, test.args)

			service := New(repo, validator, hasher, manager)

			pair, err := service.Login(test.args.ctx, test.args.username, test.args.password)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, pair)
		})
	}
}

func TestRefresh(t *testing.T) {
	type args struct {
		ctx          context.Context
		refreshToken string
	}

	tests := []struct {
		message string
		args    args
		want    model.TokenPair
		err     error
		mock    func(*mockery.PasswordValidatorMock, *mockery.PasswordHasherMock, *mockery.UserRepositoryMock, *mockery.TokenManagerMock, args)
	}{
		{
			message: "token manager error: invalid refresh token",
			args: args{
				ctx:          context.Background(),
				refreshToken: "some_refresh_token",
			},
			want: model.TokenPair{},
			err:  ErrTokenManager,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				tmm.On("ValidateRefreshToken", a.refreshToken).Once().Return(int64(0), ErrTokenManager)
				urm.AssertNotCalled(t, "FindByID", mock.Anything, mock.Anything)
				tmm.AssertNotCalled(t, "GenerateTokensPair", mock.Anything)
			},
		},
		{
			message: "user repository error: failed to find user by id",
			args: args{
				ctx:          context.Background(),
				refreshToken: "some_refresh_token",
			},
			want: model.TokenPair{},
			err:  ErrUserRepository,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				id := int64(1)

				tmm.On("ValidateRefreshToken", a.refreshToken).Once().Return(id, nil)
				urm.On("FindByID", a.ctx, id).Once().Return(model.User{}, ErrUserRepository)
				tmm.AssertNotCalled(t, "GenerateTokensPair", mock.Anything)
			},
		},
		{
			message: "token manager error: failed to generate tokens",
			args: args{
				ctx:          context.Background(),
				refreshToken: "some_refresh_token",
			},
			want: model.TokenPair{},
			err:  ErrTokenManager,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				user := model.User{ID: 1, Username: "mizentui", Password: "some_password_hash"}

				tmm.On("ValidateRefreshToken", a.refreshToken).Once().Return(user.ID, nil)
				urm.On("FindByID", a.ctx, user.ID).Once().Return(user, nil)
				tmm.On("GenerateTokensPair", user).Once().Return(model.TokenPair{}, ErrTokenManager)
			},
		},
		{
			message: "token manager ok: successfully refresh tokens",
			args: args{
				ctx:          context.Background(),
				refreshToken: "some_refresh_token",
			},
			want: model.TokenPair{AccessToken: "some_access_token", RefreshToken: "some_refresh_token"},
			err:  nil,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				user := model.User{ID: 1, Username: "mizentui", Password: "some_password_hash"}
				pair := model.TokenPair{AccessToken: "some_access_token", RefreshToken: "some_refresh_token"}

				tmm.On("ValidateRefreshToken", a.refreshToken).Once().Return(user.ID, nil)
				urm.On("FindByID", a.ctx, user.ID).Once().Return(user, nil)
				tmm.On("GenerateTokensPair", user).Once().Return(pair, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			repo := mockery.NewUserRepositoryMock(t)
			validator := mockery.NewPasswordValidatorMock(t)
			hasher := mockery.NewPasswordHasherMock(t)
			manager := mockery.NewTokenManagerMock(t)

			test.mock(validator, hasher, repo, manager, test.args)

			service := New(repo, validator, hasher, manager)

			pair, err := service.Refresh(test.args.ctx, test.args.refreshToken)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, pair)
		})
	}
}
