package client_usecase

import (
	"errors"
	"testing"
	"time"

	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/client_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestChangeAddressUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	type fields struct {
		clientRepo *mocks.ClientRepo
		timer      *mocks.Timer
	}

	type args struct {
		req UpdateClientReq
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
				req: UpdateClientReq{
					ID:          1,
					HomeAddress: "msk",
				},
			},
			want: &model.Client{
				ID:          1,
				HomeAddress: "msk",
				UpdatedAt:   now,
			},
			before: func(f fields, args args) {
				client := &model.Client{
					ID:          1,
					HomeAddress: "spb",
				}

				f.clientRepo.EXPECT().FindClient(args.req.ID).Return(client, nil)

				f.timer.EXPECT().Now().Return(now)
				client.HomeAddress = args.req.HomeAddress
				client.UpdatedAt = now

				f.clientRepo.EXPECT().UpdateClient(client).Return(client, nil)
			},
		},
		{
			name: "error on update",
			args: args{
				req: UpdateClientReq{
					ID:          1,
					HomeAddress: "msk",
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				client := &model.Client{
					ID:          1,
					HomeAddress: "spb",
				}

				f.clientRepo.EXPECT().FindClient(args.req.ID).Return(client, nil)

				f.timer.EXPECT().Now().Return(now)
				client.HomeAddress = args.req.HomeAddress
				client.UpdatedAt = now

				f.clientRepo.EXPECT().UpdateClient(client).Return(nil, errTest)
			},
		},
		{
			name: "error on find client",
			args: args{
				req: UpdateClientReq{
					ID:          1,
					HomeAddress: "msk",
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

			c, err := uc.UpdateClient(tt.args.req)

			if tt.wantErr {
				a.Error(err)

				return
			}

			a.NoError(err)
			a.Equal(tt.want, c)
		})
	}
}
