package order_usecase

import (
	"fmt"
	"time"

	"warehouse_project/internal/domain/model"
)

type AddPositionToOrderReq struct {
	order      *model.Order
	positionID int64
	quantity   int64
	unitPrice  float64
}

func (ou *OrderUseCase) AddPositionToOrder(req AddPositionToOrderReq) error {
	now := time.Now()
	position, err := ou.positionRepo.FindPosition(req.positionID)
	if err != nil {
		return fmt.Errorf("failed to find position %d: %w", req.positionID, err)
	}

	for _, orderPosition := range req.order.Positions {
		if orderPosition.PositionID == req.positionID {
			orderPosition.Quantity += req.quantity
			orderPosition.UnitPrice = float64(orderPosition.Quantity) * position.Price
			req.order.Total += float64(req.quantity) * position.Price
			req.order.UpdatedAt = time.Now()
			return nil
		}
	}

	newOrderPosition := &model.OrderPosition{
		PositionID: req.positionID,
		Quantity:   req.quantity,
		UnitPrice:  float64(req.quantity) * position.Price,
		Position:   position,
	}

	req.order.Positions = append(req.order.Positions, newOrderPosition)
	req.order.Total += newOrderPosition.UnitPrice
	req.order.UpdatedAt = now

	return nil
}
