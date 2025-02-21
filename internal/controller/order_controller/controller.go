package order_controller

import (
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/order_usecase"
)

type (
	Controller struct {
		orderUseCase orderUseCase
	}

	orderUseCase interface {
		CreateOrder(req order_usecase.CreateOrderReq) (*model.Order, error)
		FindOrder(req order_usecase.FindOrderReq) (*model.Order, error)
		UpdateOrder(req order_usecase.UpdateOrderReq) (*model.Order, error)
		DeleteOrder(req order_usecase.DeleteOrderReq) error
	}
)

func NewController(orderUseCase orderUseCase) *Controller {
	return &Controller{
		orderUseCase: orderUseCase,
	}
}
