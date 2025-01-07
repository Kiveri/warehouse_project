package position_usecase

import (
	"errors"
	"testing"

	"warehouse_project/internal/usecase/position_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type fields struct {
		positionRepo *mocks.PositionRepo
		timer        *mocks.Timer
	}

	type args struct {
		req DeletePositionReq
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
				req: DeletePositionReq{
					ID: 1,
				},
			},
			before: func(f fields, args args) {
				f.positionRepo.EXPECT().DeletePosition(args.req.ID).Return(nil)
			},
		},
		{
			name: "error on found",
			args: args{
				req: DeletePositionReq{
					ID: 50,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.positionRepo.EXPECT().DeletePosition(args.req.ID).Return(errTest)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			f := fields{
				positionRepo: mocks.NewPositionRepo(t),
				timer:        mocks.NewTimer(t),
			}

			tt.before(f, tt.args)

			uc := NewPositionUseCase(f.positionRepo, f.timer)

			err := uc.DeletePosition(tt.args.req)

			if tt.wantErr {
				a.Error(err)
				return
			}

			a.NoError(err)
		})
	}
}
