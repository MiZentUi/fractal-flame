package bcrypt

import (
	"log/slog"

	"golang.org/x/crypto/bcrypt"

	"github.com/mizentui/fractal-flame/iam/internal/config"
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
		slog.Warn("Failed to compare passwords", "err", err)

		return err
	}

	return nil
}
