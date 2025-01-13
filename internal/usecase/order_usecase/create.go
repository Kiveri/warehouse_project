package order_usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"warehouse_project/internal/domain/model"
)

var EmployeeHasNoAccessToCreateOrder = errors.New("employee has no access to create order")

type AddPositionToOrderReq struct {
	order      *model.Order
	positionID int64
	quantity   int64
	unitPrice  float64
}

func (ou *OrderUseCase) AddPositionToOrder(order *model.Order, positionID int64, quantity int64) error {
	now := time.Now()
	position, err := ou.positionRepo.FindPosition(positionID)
	if err != nil {
		return fmt.Errorf("failed to find position %d: %w", positionID, err)
	}

	for _, orderPosition := range order.Positions {
		if orderPosition.PositionID == positionID {
			orderPosition.Quantity += quantity
			orderPosition.UnitPrice = float64(orderPosition.Quantity) * position.Price
			order.Total += float64(quantity) * position.Price
			order.UpdatedAt = time.Now()
			return nil
		}
	}

	newOrderPosition := &model.OrderPosition{
		PositionID: positionID,
		Quantity:   quantity,
		UnitPrice:  float64(quantity) * position.Price,
		Position:   position,
	}

	order.Positions = append(order.Positions, newOrderPosition)
	order.Total += newOrderPosition.UnitPrice
	order.UpdatedAt = now

	return nil
}

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
		err = ou.AddPositionToOrder(order, p.PositionID, p.Quantity)
		if err != nil {
			return nil, fmt.Errorf("ou.AddPositionToOrder: %w", err)
		}
	}
	ctx := context.Background()
	order, err = ou.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("orderRepo.CreateOrder: %w", err)
	}

	return order, nil
}
