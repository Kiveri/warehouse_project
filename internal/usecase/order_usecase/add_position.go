package order_usecase

import (
	"fmt"
	"time"
	"warehouse_project/internal/domain/model"
)

type CreateOrderReq struct {
	PositionsId []int
}

func (ou *OrderUseCase) CreateOrder(req CreateOrderReq) *model.Order {
	var positions []model.Position
	for _, positionsId := range req.PositionsId {
		if position, exists := ou.positionRepo.FindPosition(positionsId); exists {
			positions = append(positions, *position)
		} else {
			fmt.Println("position does not exist")
		}
	}

	order := model.NewOrder(positions, model.OrderStatus(model.Created), model.DeliveryType(model.CourierDelivery), time.Now())
	return ou.orderRepo.CreateOrder(order)
}
