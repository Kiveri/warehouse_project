package position_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/position_usecase/mocks"
)

func TestCreateUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	// зависимости которые нужны для теста
	type fields struct {
		positionRepo *mocks.PositionRepo
		timer        *mocks.Timer
	}
	// данные для теста
	type args struct {
		req CreatePositionReq
	}
	// тесты
	tests := []struct {
		name    string
		args    args
		want    *model.Position
		wantErr bool                      // если хотим ошибку
		before  func(f fields, args args) // замокать наши вызовы перед тестом
	}{
		{
			name: "success",
			args: args{
				req: CreatePositionReq{
					Name:    "Йогурт",
					Barcode: "1234",
					Price:   10000,
					PosType: model.Liquid,
				},
			},
			want: &model.Position{
				Name:      "Йогурт",
				Barcode:   "1234",
				Price:     10000,
				PosType:   model.Liquid,
				CreatedAt: now,
				UpdatedAt: now,
			},
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				position := model.NewPosition(args.req.Name, args.req.Barcode, args.req.Price, args.req.PosType, now)
				f.positionRepo.EXPECT().CreatePosition(position).Return(position, nil)
			},
		},
		{
			name: "error on create",
			args: args{
				req: CreatePositionReq{
					Name:    "Йогурт",
					Barcode: "1234",
					Price:   10000,
					PosType: model.Liquid,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				position := model.NewPosition(args.req.Name, args.req.Barcode, args.req.Price, args.req.PosType, now)
				f.positionRepo.EXPECT().CreatePosition(position).Return(nil, errTest)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			// создали зависимости
			f := fields{
				positionRepo: mocks.NewPositionRepo(t),
				timer:        mocks.NewTimer(t),
			}
			tt.before(f, tt.args)

			uc := NewPositionUseCase(f.positionRepo, f.timer)

			// выполнили
			position, err := uc.CreatePosition(tt.args.req)

			// проверяем результат
			if tt.wantErr {
				a.Error(err)

				return
			}
			a.NoError(err)
			a.Equal(tt.want, position)
		})
	}
}
