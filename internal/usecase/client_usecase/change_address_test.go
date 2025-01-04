package client_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/client_usecase/mocks"
)

func TestChangeAddressUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now()
	//errTest := errors.New("test error")

	type fields struct {
		clientRepo *mocks.ClientRepo
		timer      *mocks.Timer
	}

	type args struct {
		req UpdateClientReq
	}

	client := &model.Client{
		ID:          1,
		HomeAddress: "spb",
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
					HomeAddress: "spb",
				},
			},
			want: &model.Client{
				ID:          1,
				HomeAddress: "sbp",
				UpdatedAt:   now,
			},
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				updatedClient := &model.Client{
					ID:          args.req.ID,
					HomeAddress: args.req.HomeAddress,
					UpdatedAt:   now,
				}
				f.clientRepo.EXPECT().UpdateClient(client).Return(client, nil)
				f.clientRepo.EXPECT().FindClient(args.req.ID).Return(updatedClient, nil)
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
			a.NotEqual(tt.want, c)
		})
	}
}
