package order_usecase

import (
	"errors"
	"fmt"
	"time"
	"warehouse_project/internal/domain/model"
)

var EmployeeHasNoAccessToCreateOrder = errors.New("employee has no access to create order")

type CreateOrderReq struct {
	EmployeeID   int64
	PositionsIDs []int64
	DeliveryType model.DeliveryType
}

func (ou *OrderUseCase) CreateOrder(req CreateOrderReq) (*model.Order, error) {
	employee, err := ou.employeeRepo.FindEmployee(req.EmployeeID)
	if err != nil {
		return nil, fmt.Errorf("employeeRepo.FindEmployee: %w", err)
	}

	if !employee.IsCanOrderCreate() {
		return nil, EmployeeHasNoAccessToCreateOrder
	}

	positions, err := ou.positionRepo.FindAllByIDs(req.PositionsIDs)
	if err != nil {
		return nil, fmt.Errorf("positionRepo.FindAllByIDs: %w", err)
	}

	order := model.NewOrder(positions, employee.ID, req.DeliveryType, time.Now())
	order, err = ou.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, fmt.Errorf("orderRepo.CreateOrder: %w", err)
	}

	return order, nil
}
