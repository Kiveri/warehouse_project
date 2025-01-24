package order_usecase

import (
	"errors"
	"testing"
	"time"

	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/order_usecase/mocks"

	"github.com/stretchr/testify/assert"
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

	employeeWithAccess := &model.Employee{
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

	position1 := &model.Position{
		ID:    1,
		Name:  "Name1",
		Price: 100,
	}
	position2 := &model.Position{
		ID:    2,
		Name:  "Name2",
		Price: 200,
	}

	orderPosition1 := &model.OrderPosition{
		Quantity:  1,
		UnitPrice: 100,
		Position:  position1,
	}
	orderPosition2 := &model.OrderPosition{
		Quantity:  1,
		UnitPrice: 200,
		Position:  position2,
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
					PositionIDs:  []int64{position1.ID, position2.ID},
					EmployeeID:   employeeWithAccess.ID,
					ClientID:     1,
					Status:       model.Created,
					DeliveryType: model.PointOfDelivery,
				},
			},
			want: &model.Order{
				Positions: map[int64]*model.OrderPosition{
					position1.ID: orderPosition1,
					position2.ID: orderPosition2,
				},
				EmployeeID:   employeeWithAccess.ID,
				ClientID:     1,
				Status:       model.Created,
				DeliveryType: model.PointOfDelivery,
				Total:        0,
				CreatedAt:    now,
			},
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				f.employeeRepo.EXPECT().FindEmployee(args.req.EmployeeID).Return(employeeWithAccess, nil)
				f.clientRepo.EXPECT().FindClient(args.req.ClientID).Return(client, nil)
				f.positionRepo.EXPECT().FindPosition(position1.ID).Return(position1, nil)
				f.positionRepo.EXPECT().FindPosition(position2.ID).Return(position2, nil)
				order := model.NewOrder(employeeWithAccess.ID, client.ID, args.req.DeliveryType, now)
				order.AddPositions([]*model.Position{
					position1,
					position2,
				})
				f.orderRepo.EXPECT().CreateOrder(order).Return(order, nil)
			},
		},
		{
			name: "error on create",
			args: args{
				req: CreateOrderReq{
					PositionIDs:  []int64{position1.ID, position2.ID},
					EmployeeID:   employeeWithAccess.ID,
					ClientID:     1,
					Status:       model.Created,
					DeliveryType: model.PointOfDelivery,
				},
			},
			wantErr: errTest,
			before: func(f fields, args args) {
				f.timer.EXPECT().Now().Return(now)
				f.employeeRepo.EXPECT().FindEmployee(args.req.EmployeeID).Return(employeeWithAccess, nil)
				f.clientRepo.EXPECT().FindClient(args.req.ClientID).Return(client, nil)
				f.positionRepo.EXPECT().FindPosition(position1.ID).Return(position1, nil)
				f.positionRepo.EXPECT().FindPosition(position2.ID).Return(position2, nil)
				order := model.NewOrder(employeeWithAccess.ID, client.ID, args.req.DeliveryType, now)
				order.AddPositions([]*model.Position{
					position1,
					position2,
				})
				f.orderRepo.EXPECT().CreateOrder(order).Return(nil, errTest)
			},
		},
		{
			name: "error no access",
			args: args{
				req: CreateOrderReq{
					PositionIDs:  []int64{position1.ID, position2.ID},
					EmployeeID:   employeeWithoutAccess.ID,
					ClientID:     1,
					Status:       model.Created,
					DeliveryType: model.PointOfDelivery,
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
