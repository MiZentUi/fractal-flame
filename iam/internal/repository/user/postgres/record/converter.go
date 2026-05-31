package record

import "github.com/mizentui/fractal-flame/iam/internal/model"

func RawToModel(user UserRow) model.User {
	var image string
	if user.Image.Valid {
		image = user.Image.String
	}

	return model.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Image:    image,
	}
}
