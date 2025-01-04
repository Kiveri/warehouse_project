package client_usecase

import (
	"errors"
	"testing"
	"time"

	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/client_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	type fields struct {
		clientRepo *mocks.ClientRepo
		timer      *mocks.Timer
	}

	type args struct {
		req CreateClientReq
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
				req: CreateClientReq{
					Name:        "Имя",
					Phone:       "899999999",
					Email:       "test@test.com",
					HomeAddress: "spb",
				},
			},
			want: &model.Client{
				Name:        "Имя",
				Phone:       "899999999",
				Email:       "test@test.com",
				HomeAddress: "spb",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				client := model.NewClient(args.req.Name, args.req.Phone, args.req.Email, args.req.HomeAddress, now)
				f.clientRepo.EXPECT().CreateClient(client).Return(client, nil)
			},
		},
		{
			name: "error on create",
			args: args{
				req: CreateClientReq{
					Name:        "Имя",
					Phone:       "899999999",
					Email:       "test@test.com",
					HomeAddress: "spb",
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				client := model.NewClient(args.req.Name, args.req.Phone, args.req.Email, args.req.HomeAddress, now)
				f.clientRepo.EXPECT().CreateClient(client).Return(nil, errTest)
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

			client, err := uc.CreateClient(tt.args.req)

			if tt.wantErr {
				a.Error(err)

				return
			}
			a.NoError(err)
			a.Equal(tt.want, client)
		})
	}
}
