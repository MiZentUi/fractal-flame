package image

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	mockery "github.com/mizentui/fractal-flame/iam/internal/service/image/mock"
)

var ErrImageRepository = errors.New("some image repo error")

func TestGetImage(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}

	tests := []struct {
		message string
		args    args
		want    []byte
		err     error
		mock    func(*mockery.ImageRepositoryMock, args)
	}{
		{
			message: "image repository error: failed to find image",
			args: args{
				ctx:  context.Background(),
				name: "avatar.png",
			},
			want: nil,
			err:  ErrImageRepository,
			mock: func(irm *mockery.ImageRepositoryMock, a args) {
				irm.On("FindByName", a.ctx, a.name).Once().Return(nil, ErrImageRepository)
			},
		},
		{
			message: "image repository ok: successfully find image",
			args: args{
				ctx:  context.Background(),
				name: "avatar.png",
			},
			want: []byte("some image bytes"),
			err:  nil,
			mock: func(irm *mockery.ImageRepositoryMock, a args) {
				image := []byte("some image bytes")

				irm.On("FindByName", a.ctx, a.name).Once().Return(image, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			t.Parallel()

			repo := mockery.NewImageRepositoryMock(t)

			test.mock(repo, test.args)

			service := New(repo)

			image, err := service.GetImage(test.args.ctx, test.args.name)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, image)
		})
	}
}
