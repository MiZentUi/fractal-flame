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

const (
	userID                      int64 = 1
	emptyUserID                 int64 = 0
	username                          = "mizentui"
	password                          = "1238124AAs"
	shortPassword                     = "123"
	noSpecialCharactersPassword       = "123123123123123123"
	passwordHash                      = "some_password_hash"
	accessToken                       = "some_access_token"
	refreshToken                      = "some_refresh_token"
)

var (
	ErrPasswordHasher = errors.New("some hasher error")
	ErrUserRepository = errors.New("some user repo error")
	ErrTokenManager   = errors.New("some token manager error")
	ErrShortPassword  = errors.New("password is too short")
	ErrSpecialChars   = errors.New("require special characters in password")
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
				username: username,
				password: shortPassword,
			},
			want: emptyUserID,
			err:  errs.ErrWeakPassword,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				pvm.On("Validate", a.password).Once().Return(ErrShortPassword)
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				urm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "validator error: require special characters",
			args: args{
				ctx:      context.Background(),
				username: username,
				password: noSpecialCharactersPassword,
			},
			want: emptyUserID,
			err:  errs.ErrWeakPassword,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				pvm.On("Validate", a.password).Once().Return(ErrSpecialChars)
				phm.AssertNotCalled(t, "HashAndSalt", mock.Anything)
				urm.AssertNotCalled(t, "Save", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			message: "password hasher error: failed to get hash",
			args: args{
				ctx:      context.Background(),
				username: username,
				password: password,
			},
			want: emptyUserID,
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
				username: username,
				password: password,
			},
			want: emptyUserID,
			err:  ErrUserRepository,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(passwordHash, nil)
				urm.On("Save", a.ctx, a.username, passwordHash).Once().Return(emptyUserID, ErrUserRepository)
			},
		},
		{
			message: "user repository ok: successfully save user",
			args: args{
				ctx:      context.Background(),
				username: username,
				password: password,
			},
			want: userID,
			err:  nil,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				pvm.On("Validate", a.password).Once().Return(nil)
				phm.On("HashAndSalt", a.password).Once().Return(passwordHash, nil)
				urm.On("Save", a.ctx, a.username, passwordHash).Once().Return(userID, nil)
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
				username: username,
				password: password,
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
				username: username,
				password: password,
			},
			want: model.TokenPair{},
			err:  errs.ErrInvalidCredentials,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				user := model.User{ID: userID, Username: a.username, Password: passwordHash}

				urm.On("FindByUsername", a.ctx, a.username).Once().Return(user, nil)
				phm.On("ComparePasswords", user.Password, a.password).Once().Return(ErrPasswordHasher)
				tmm.AssertNotCalled(t, "GenerateTokensPair", mock.Anything)
			},
		},
		{
			message: "token manager error: failed to generate tokens",
			args: args{
				ctx:      context.Background(),
				username: username,
				password: password,
			},
			want: model.TokenPair{},
			err:  ErrTokenManager,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				user := model.User{ID: userID, Username: a.username, Password: passwordHash}

				urm.On("FindByUsername", a.ctx, a.username).Once().Return(user, nil)
				phm.On("ComparePasswords", user.Password, a.password).Once().Return(nil)
				tmm.On("GenerateTokensPair", user).Once().Return(model.TokenPair{}, ErrTokenManager)
			},
		},
		{
			message: "token manager ok: successfully login",
			args: args{
				ctx:      context.Background(),
				username: username,
				password: password,
			},
			want: model.TokenPair{AccessToken: accessToken, RefreshToken: refreshToken},
			err:  nil,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				user := model.User{ID: userID, Username: a.username, Password: passwordHash}
				pair := model.TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}

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
				refreshToken: refreshToken,
			},
			want: model.TokenPair{},
			err:  ErrTokenManager,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				tmm.On("ValidateRefreshToken", a.refreshToken).Once().Return(emptyUserID, ErrTokenManager)
				urm.AssertNotCalled(t, "FindByID", mock.Anything, mock.Anything)
				tmm.AssertNotCalled(t, "GenerateTokensPair", mock.Anything)
			},
		},
		{
			message: "user repository error: failed to find user by id",
			args: args{
				ctx:          context.Background(),
				refreshToken: refreshToken,
			},
			want: model.TokenPair{},
			err:  ErrUserRepository,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				tmm.On("ValidateRefreshToken", a.refreshToken).Once().Return(userID, nil)
				urm.On("FindByID", a.ctx, userID).Once().Return(model.User{}, ErrUserRepository)
				tmm.AssertNotCalled(t, "GenerateTokensPair", mock.Anything)
			},
		},
		{
			message: "token manager error: failed to generate tokens",
			args: args{
				ctx:          context.Background(),
				refreshToken: refreshToken,
			},
			want: model.TokenPair{},
			err:  ErrTokenManager,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				user := model.User{ID: userID, Username: username, Password: passwordHash}

				tmm.On("ValidateRefreshToken", a.refreshToken).Once().Return(user.ID, nil)
				urm.On("FindByID", a.ctx, user.ID).Once().Return(user, nil)
				tmm.On("GenerateTokensPair", user).Once().Return(model.TokenPair{}, ErrTokenManager)
			},
		},
		{
			message: "token manager ok: successfully refresh tokens",
			args: args{
				ctx:          context.Background(),
				refreshToken: refreshToken,
			},
			want: model.TokenPair{AccessToken: accessToken, RefreshToken: refreshToken},
			err:  nil,
			mock: func(pvm *mockery.PasswordValidatorMock, phm *mockery.PasswordHasherMock, urm *mockery.UserRepositoryMock, tmm *mockery.TokenManagerMock, a args) {
				user := model.User{ID: userID, Username: username, Password: passwordHash}
				pair := model.TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}

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
