package order_usecase

import (
	"time"

	"warehouse_project/internal/domain/model"
)

type OrderUseCase struct {
	positionRepo positionRepo
	orderRepo    orderRepo
	employeeRepo employeeRepo
	clientRepo   clientRepo
	timer        timer
}

type (
	orderRepo interface {
		CreateOrder(order *model.Order) (*model.Order, error)
		FindOrder(id int64) (*model.Order, error)
		UpdateOrder(order *model.Order) (*model.Order, error)
		DeleteOrder(id int64) error
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

func NewOrderUseCase(orderRepo orderRepo, positionRepo positionRepo, employeeRepo employeeRepo, clientRepo clientRepo, timer timer) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:    orderRepo,
		positionRepo: positionRepo,
		employeeRepo: employeeRepo,
		clientRepo:   clientRepo,
		timer:        timer,
	}
}
