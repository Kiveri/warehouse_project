package position_usecase

import (
	"errors"
	"testing"
	"time"

	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/position_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestChangePriceUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	type fields struct {
		positionRepo *mocks.PositionRepo
		timer        *mocks.Timer
	}

	type args struct {
		req UpdatePositionReq
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
				req: UpdatePositionReq{
					ID:    1,
					Price: 123,
				},
			},
			want: &model.Position{
				ID:        1,
				Price:     123,
				UpdatedAt: now,
			},
			before: func(f fields, args args) {
				position := &model.Position{
					ID:    1,
					Price: 124,
				}

				f.positionRepo.EXPECT().FindPosition(args.req.ID).Return(position, nil)

				f.timer.EXPECT().Now().Return(now)
				position.Price = args.req.Price
				position.UpdatedAt = now

				f.positionRepo.EXPECT().UpdatePosition(position).Return(position, nil)
			},
		},
		{
			name: "error on update",
			args: args{
				req: UpdatePositionReq{
					ID:    1,
					Price: 123,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				position := &model.Position{
					ID:    1,
					Price: 124,
				}

				f.positionRepo.EXPECT().FindPosition(args.req.ID).Return(position, nil)

				f.timer.EXPECT().Now().Return(now)
				position.Price = args.req.Price
				position.UpdatedAt = now

				f.positionRepo.EXPECT().UpdatePosition(position).Return(nil, errTest)
			},
		},
		{
			name: "error on find",
			args: args{
				req: UpdatePositionReq{
					ID:    1,
					Price: 123,
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

			p, err := uc.UpdatePosition(tt.args.req)

			if tt.wantErr {
				a.Error(err)

				return
			}

			a.NoError(err)
			a.Equal(tt.want, p)
		})
	}
}
