package order_usecase

import (
	"errors"
	"testing"

	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/order_usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestFindUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	order := &model.Order{
		ID: 1,
	}

	type fields struct {
		orderRepo    *mocks.OrderRepo
		positionRepo *mocks.PositionRepo
		employeeRepo *mocks.EmployeeRepo
		clientRepo   *mocks.ClientRepo
		timer        *mocks.Timer
	}

	type args struct {
		req FindOrderReq
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
				req: FindOrderReq{
					ID: 1,
				},
			},
			want: &model.Order{
				ID: 1,
			},
			before: func(f fields, args args) {
				f.orderRepo.EXPECT().FindOrder(args.req.ID).Return(order, nil)
			},
		},
		{
			name: "error on find",
			args: args{
				req: FindOrderReq{
					ID: 2,
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

			c, err := uc.FindOrder(tt.args.req)

			if tt.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(tt.want, c)
		})
	}
}
