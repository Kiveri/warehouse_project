package order_usecase

import (
	"errors"
	"fmt"

	"warehouse_project/internal/domain/model"
)

var EmployeeHasNoAccessToCreateOrder = errors.New("employee has no access to create order")

type CreateOrderReq struct {
	Positions    []*model.OrderPosition
	EmployeeID   int64
	ClientID     int64
	Status       model.OrderStatus
	DeliveryType model.DeliveryType
}

func (ou *OrderUseCase) CreateOrder(req CreateOrderReq) (*model.Order, error) {
	now := ou.timer.Now()

	employee, err := ou.employeeRepo.FindEmployee(req.EmployeeID)
	if err != nil {

		return nil, fmt.Errorf("employeeRepo.FindEmployee: %w", err)
	}

	if !employee.IsCanOrderCreate() {

		return nil, EmployeeHasNoAccessToCreateOrder
	}

	client, err := ou.clientRepo.FindClient(req.ClientID)
	if err != nil {

		return nil, fmt.Errorf("ou.ClientRepo.FindClient: %w", err)
	}

	order := model.NewOrder(employee.ID, client.ID, req.DeliveryType, now)
	for _, p := range req.Positions {
		err = ou.AddPositionToOrder(AddPositionToOrderReq{
			orderID:    order.ID,
			positionID: p.PositionID,
			quantity:   p.Quantity,
		})
		if err != nil {

			return nil, fmt.Errorf("ou.AddPositionToOrder: %w", err)
		}
	}

	order, err = ou.orderRepo.CreateOrder(order)
	if err != nil {

		return nil, fmt.Errorf("orderRepo.CreateOrder: %w", err)
	}

	return order, nil
}
