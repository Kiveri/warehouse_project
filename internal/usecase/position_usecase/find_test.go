package position_usecase

import (
	"errors"
	"testing"

	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/position_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestFindUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	position := &model.Position{
		ID: 1,
	}

	type fields struct {
		positionRepo *mocks.PositionRepo
		timer        *mocks.Timer
	}

	type args struct {
		req FindPositionReq
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Position
		wantErr bool
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				req: FindPositionReq{
					ID: 1,
				},
			},
			want: &model.Position{
				ID: 1,
			},
			before: func(f fields, args args) {
				f.positionRepo.EXPECT().FindPosition(args.req.ID).Return(position, nil)
			},
		},
		{
			name: "error on find",
			args: args{
				req: FindPositionReq{
					ID: 2,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.positionRepo.EXPECT().FindPosition(args.req.ID).Return(nil, errTest)
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

			c, err := uc.FindPosition(tt.args.req)

			if tt.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(tt.want, c)
		})
	}
}
