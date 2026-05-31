package jwt

import (
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	errs "github.com/mizentui/fractal-flame/iam/internal/error"
	"github.com/mizentui/fractal-flame/iam/internal/model"
)

type manager struct {
	accessSigningKey  string
	refreshSigningKey string
	accessTokenTTL    time.Duration
	refreshTokenTTL   time.Duration
}

func New(accessSigningKey, refreshSigningKey string, accessTokenTTL, refreshTokenTTL time.Duration) *manager {
	return &manager{
		accessSigningKey:  accessSigningKey,
		refreshSigningKey: refreshSigningKey,
		accessTokenTTL:    accessTokenTTL,
		refreshTokenTTL:   refreshTokenTTL,
	}
}

const (
	accessTokenType  string = "access"
	refreshTokenType string = "refresh"
)

type Claims struct {
	Username string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`

	jwt.RegisteredClaims
}

func (m *manager) GenerateTokensPair(user model.User) (model.TokenPair, error) {
	accessToken, err := m.generateToken(user, accessTokenType, m.accessTokenTTL, m.accessSigningKey)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("generate access token: %w", err)
	}

	refreshToken, err := m.generateToken(user, refreshTokenType, m.refreshTokenTTL, m.refreshSigningKey)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("generate refresh token: %w", err)
	}

	slog.Debug("Successfully generate tokens")

	return model.TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (m *manager) ValidateRefreshToken(token string) (int64, error) {
	return m.validateToken(token, m.refreshSigningKey, refreshTokenType)
}

func (m *manager) ValidateAccessToken(token string) (int64, error) {
	return m.validateToken(token, m.accessSigningKey, accessTokenType)
}

func (m *manager) validateToken(tokenString string, signingKey string, tokenType string) (int64, error) {
	var claims Claims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS384 {
			slog.Warn("Failed, invalid signing method", "method", t.Method)

			return nil, errs.ErrInvalidToken
		}

		return []byte(signingKey), nil
	})

	if err != nil || !token.Valid {
		slog.Error("Failed, invalid token")

		return 0, errs.ErrInvalidToken
	}

	if claims.Type != tokenType {
		slog.Error("Failed, invalid token type", "type", claims.Type)

		return 0, errs.ErrInvalidToken
	}

	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		slog.Error("Failed to convert string to int", "id", id, "err", err)

		return 0, errs.ErrInvalidToken
	}

	return int64(id), err
}

func (m *manager) generateToken(user model.User, tokenType string, tokenTTL time.Duration, signingKey string) (string, error) {
	claims := Claims{
		Username: user.Username,
		Type:     tokenType,

		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprint(user.ID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		slog.Error("Failed to sign token", "type", tokenType, "err", err)

		return "", err
	}

	return tokenString, nil
}
