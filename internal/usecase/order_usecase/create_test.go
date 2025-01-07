package order_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/order_usecase/mocks"
)

func TestCreateUseCase(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	type fields struct {
		orderRepo    *mocks.OrderRepo
		clientRepo   *mocks.ClientRepo
		positionRepo *mocks.PositionRepo
		employeeRepo *mocks.EmployeeRepo
		timer        *mocks.Timer
	}

	type args struct {
		req CreateOrderReq
	}

	employee := &model.Employee{
		ID:   1,
		Role: model.Manager,
	}

	employeeWithoutAccess := &model.Employee{
		ID:   2,
		Role: model.Specialist,
	}

	client := &model.Client{
		ID: 1,
	}

	positions := []*model.Position{
		{
			ID:    2,
			Name:  "Позиция1",
			Price: 1234,
		},
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Order
		wantErr error
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				req: CreateOrderReq{
					EmployeeID:   1,
					PositionsIDs: []int64{2},
					DeliveryType: model.CourierDelivery,
					ClientID:     1,
				},
			},
			want: &model.Order{
				Positions: []*model.Position{
					{
						ID:    2,
						Name:  "Позиция1",
						Price: 1234,
					},
				},
				CreatedBy: 1,
				Client:    1,
				Status:    model.Created,
				DelType:   model.CourierDelivery,
				Total:     1234,
				CreatedAt: now,
			},
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				f.employeeRepo.EXPECT().FindEmployee(args.req.EmployeeID).Return(employee, nil)
				f.clientRepo.EXPECT().FindClient(args.req.ClientID).Return(client, nil)
				f.positionRepo.EXPECT().FindAllByIDs(args.req.PositionsIDs).Return(positions, nil)
				order := model.NewOrder(positions, employee.ID, client.ID, args.req.DeliveryType, now)
				f.orderRepo.EXPECT().CreateOrder(order).Return(order, nil)
			},
		},
		{
			name: "error on create",
			args: args{
				req: CreateOrderReq{
					EmployeeID:   1,
					PositionsIDs: []int64{2},
					DeliveryType: model.CourierDelivery,
					ClientID:     1,
				},
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				f.employeeRepo.EXPECT().FindEmployee(args.req.EmployeeID).Return(employee, nil)
				f.clientRepo.EXPECT().FindClient(args.req.ClientID).Return(client, nil)
				f.positionRepo.EXPECT().FindAllByIDs(args.req.PositionsIDs).Return(positions, nil)
				order := model.NewOrder(positions, employee.ID, client.ID, args.req.DeliveryType, now)
				f.orderRepo.EXPECT().CreateOrder(order).Return(nil, errTest)
			},
		},
		{
			name: "error no access",
			args: args{
				req: CreateOrderReq{
					EmployeeID:   2,
					PositionsIDs: []int64{2},
					DeliveryType: model.CourierDelivery,
					ClientID:     1,
				},
			},
			wantErr: EmployeeHasNoAccessToCreateOrder,
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				f.employeeRepo.EXPECT().FindEmployee(args.req.EmployeeID).Return(employeeWithoutAccess, nil)
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
				clientRepo:   mocks.NewClientRepo(t),
				employeeRepo: mocks.NewEmployeeRepo(t),
				orderRepo:    mocks.NewOrderRepo(t),
				timer:        mocks.NewTimer(t),
			}
			tt.before(f, tt.args)

			uc := NewOrderUseCase(f.orderRepo, f.positionRepo, f.employeeRepo, f.clientRepo, f.timer)

			// выполнили
			order, err := uc.CreateOrder(tt.args.req)

			// проверяем результат
			if tt.wantErr != nil {
				a.ErrorIs(err, tt.wantErr)

				return
			}
			a.NoError(err)
			a.Equal(tt.want, order)
		})
	}
}
