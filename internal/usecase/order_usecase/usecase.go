package order_usecase

import (
	"time"
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
	clientRepo interface {
		FindClient(id int64) (*model.Client, error)
	}
	timer interface {
		Now() time.Time
	}
)

type OrderUseCase struct {
	positionRepo positionRepo
	orderRepo    orderRepo
	employeeRepo employeeRepo
	clientRepo   clientRepo
	timer        timer
}

func NewOrderUseCase(orderRepo orderRepo, positionRepo positionRepo, employeeRepo employeeRepo, clientRepo clientRepo, timer timer) *OrderUseCase {
	return &OrderUseCase{
		positionRepo: positionRepo,
		orderRepo:    orderRepo,
		employeeRepo: employeeRepo,
		clientRepo:   clientRepo,
		timer:        timer,
	}
}
