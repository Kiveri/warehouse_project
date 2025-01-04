package order_usecase

import (
	"errors"
	"testing"
	"time"

	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/order_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestChangeStatusUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	type fields struct {
		orderRepo    *mocks.OrderRepo
		positionRepo *mocks.PositionRepo
		employeeRepo *mocks.EmployeeRepo
		clientRepo   *mocks.ClientRepo
		timer        *mocks.Timer
	}

	type args struct {
		req UpdateOrderReq
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Order
		wantErr bool
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				req: UpdateOrderReq{
					ID:     1,
					Status: model.Building,
				},
			},
			want: &model.Order{
				ID:        1,
				Status:    model.Building,
				UpdatedAt: now,
			},
			before: func(f fields, args args) {
				order := &model.Order{
					ID:     1,
					Status: model.Created,
				}

				f.orderRepo.EXPECT().FindOrder(args.req.ID).Return(order, nil)

				f.timer.EXPECT().Now().Return(now)
				order.Status = args.req.Status
				order.UpdatedAt = now

				f.orderRepo.EXPECT().UpdateOrder(order).Return(order, nil)
			},
		},
		{
			name: "error on update",
			args: args{
				req: UpdateOrderReq{
					ID:     1,
					Status: model.Building,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				order := &model.Order{
					ID:     1,
					Status: model.Created,
				}

				f.orderRepo.EXPECT().FindOrder(args.req.ID).Return(order, nil)

				f.timer.EXPECT().Now().Return(now)
				order.Status = args.req.Status
				order.UpdatedAt = now

				f.orderRepo.EXPECT().UpdateOrder(order).Return(nil, errTest)
			},
		},
		{
			name: "error on find",
			args: args{
				req: UpdateOrderReq{
					ID:     1,
					Status: model.Building,
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.orderRepo.EXPECT().FindOrder(args.req.ID).Return(nil, errTest)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			f := fields{
				orderRepo:    mocks.NewOrderRepo(t),
				positionRepo: mocks.NewPositionRepo(t),
				employeeRepo: mocks.NewEmployeeRepo(t),
				clientRepo:   mocks.NewClientRepo(t),
				timer:        mocks.NewTimer(t),
			}
			tt.before(f, tt.args)

			uc := NewOrderUseCase(f.orderRepo, f.positionRepo, f.employeeRepo, f.clientRepo, f.timer)

			o, err := uc.UpdateOrder(tt.args.req)

			if tt.wantErr {
				a.Error(err)

				return
			}

			a.NoError(err)
			a.Equal(tt.want, o)
		})
	}
}
