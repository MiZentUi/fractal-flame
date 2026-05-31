package dto

import (
	"github.com/mizentui/fractal-flame/iam/internal/model"
	iamv1 "github.com/mizentui/fractal-flame/iam/pkg/proto/v1"
)

func ModelToProto(user model.User) *iamv1.User {
	return &iamv1.User{
		UserId:   user.ID,
		Username: user.Username,
		Image:    user.Image,
	}
}
