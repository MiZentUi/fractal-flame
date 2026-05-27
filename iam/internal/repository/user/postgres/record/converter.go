package record

import "github.com/mizentui/fractal-flame/iam/internal/model"

func RawToModel(user UserRow) model.User {
	return model.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Image:    user.Image,
	}
}
