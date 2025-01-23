package order_usecase

import (
	"errors"
	"fmt"

	"warehouse_project/internal/domain/model"
)

var EmployeeHasNoAccessToCreateOrder = errors.New("employee has no access to create order")

// клиент выбирает позиции
// идентификаторы позиций на создание заказа
// + написать юзкейс списка позиций
// + список заказов клиента (должна возвращаться шорт информация)

type CreateOrderReq struct {
	PositionIDs  []int64
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

	positions := make([]*model.Position, 0, len(req.PositionIDs))
	for _, positionID := range req.PositionIDs {
		var position *model.Position
		position, err = ou.positionRepo.FindPosition(positionID)
		if err != nil {
			return nil, fmt.Errorf("ou.positionRepo.FindPosition: %w", err)
		}

		positions = append(positions, position)
	}

	order := model.NewOrder(employee.ID, client.ID, req.DeliveryType, now)
	order.AddPositions(positions)

	order, err = ou.orderRepo.CreateOrder(order)
	if err != nil {

		return nil, fmt.Errorf("orderRepo.CreateOrder: %w", err)
	}

	return order, nil
}
