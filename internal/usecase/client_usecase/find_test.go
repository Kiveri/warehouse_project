package client_usecase

import (
	"errors"
	"testing"

	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/client_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestFindUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	client := &model.Client{
		ID: 1,
	}

	type fields struct {
		clientRepo *mocks.ClientRepo
		timer      *mocks.Timer
	}

	type args struct {
		req FindClientReq
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Client
		wantErr bool
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				req: FindClientReq{
					ID: 1,
				},
			},
			want: &model.Client{
				ID: 1,
			},
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindClient(args.req.ID).Return(client, nil)
			},
		},
		{
			name: "error on find",
			args: args{
				req: FindClientReq{
					ID: 2,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindClient(args.req.ID).Return(nil, errTest)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			f := fields{
				clientRepo: mocks.NewClientRepo(t),
				timer:      mocks.NewTimer(t),
			}

			tt.before(f, tt.args)

			uc := NewClientUseCase(f.clientRepo, f.timer)

			c, err := uc.FindClient(tt.args.req)

			if tt.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(tt.want, c)
		})
	}
}
