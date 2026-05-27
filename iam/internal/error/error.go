package error

import "errors"

var (
	ErrWeakPassword       error = errors.New("weak password")
	ErrUserAlreadyExists  error = errors.New("user with this username already exists")
	ErrInvalidCredentials error = errors.New("invalid credentials")
	ErrInvalidToken       error = errors.New("invalid token")
	ErrInvalidCtxValue    error = errors.New("invalid ctx value")
	ErrInvalidUserID      error = errors.New("invalid user id")
	ErrNothingToUpdate    error = errors.New("nothing to update")
	ErrInvalidImageName   error = errors.New("invalid image name")
	ErrUserNotFound       error = errors.New("user not found")
)
