package client_usecase

import (
	"errors"
	"testing"

	"warehouse_project/internal/usecase/client_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type fields struct {
		clientRepo *mocks.ClientRepo
		timer      *mocks.Timer
	}

	type args struct {
		req DeleteClientReq
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				req: DeleteClientReq{
					ID: 1,
				},
			},
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().DeleteClient(args.req.ID).Return(nil)
			},
		},
		{
			name: "client not found",
			args: args{
				req: DeleteClientReq{
					ID: 50,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.clientRepo.EXPECT().DeleteClient(args.req.ID).Return(errTest)
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

			err := uc.DeleteClient(tt.args.req)

			if tt.wantErr {
				a.Error(err)
				return
			}

			a.NoError(err)
		})
	}
}
