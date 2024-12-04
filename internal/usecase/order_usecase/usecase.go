package order_usecase

import (
	"warehouse_project/internal/domain/model"
)

type (
	orderRepo interface {
		CreateOrder(order *model.Order) (*model.Order, error)
	}
	positionRepo interface {
		FindAllByIDs(id []int64) ([]*model.Position, error)
	}
	employeeRepo interface {
		FindEmployee(id int64) (*model.Employee, error)
	}
)

type OrderUseCase struct {
	positionRepo positionRepo
	orderRepo    orderRepo
	employeeRepo employeeRepo
}

func NewOrderUseCase(orderRepo orderRepo, positionRepo positionRepo, employeeRepo employeeRepo) *OrderUseCase {
	return &OrderUseCase{
		positionRepo: positionRepo,
		orderRepo:    orderRepo,
		employeeRepo: employeeRepo,
	}
}
