package client_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/client_usecase/mocks"
)

func TestFindUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")
	errDb := errors.New("database error")

	client := &model.Client{
		ID:          1,
		Name:        "Name",
		HomeAddress: "spb",
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
				ID:          1,
				Name:        "Name",
				HomeAddress: "spb",
			},
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindClient(args.req.ID).Return(client, nil)
			},
		},
		{
			name: "client not found",
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
		{
			name: "db error",
			args: args{
				req: FindClientReq{
					ID: 1,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().FindClient(args.req.ID).Return(nil, errDb)
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
