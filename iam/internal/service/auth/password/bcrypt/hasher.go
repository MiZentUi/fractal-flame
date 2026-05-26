package bcrypt

import (
	"log/slog"

	"github.com/mizentui/fractal-flame/iam/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type hasher struct{}

func New() *hasher {
	return &hasher{}
}

func (h *hasher) HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), config.App().Auth.BcryptCost())
	if err != nil {
		slog.Error("Failed to get password hash", "err", err)

		return "", err
	}

	return string(hash), nil
}

func (h *hasher) ComparePasswords(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		slog.Warn("Failed, invalid credentials", "err", err)

		return err
	}

	return nil
}
