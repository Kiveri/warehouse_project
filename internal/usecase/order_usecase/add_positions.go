package order_usecase

import (
	"fmt"
)

type AddPositionToOrderReq struct {
	orderID    int64
	positionID int64
	quantity   int64
}

func (ou *OrderUseCase) AddPositionToOrder(req AddPositionToOrderReq) error {
	order, err := ou.orderRepo.FindOrder(req.orderID)
	if err != nil {

		return fmt.Errorf("order with id %d not found", req.orderID)
	}
	position, err := ou.positionRepo.FindPosition(req.positionID)
	if err != nil {

		return fmt.Errorf("position with id %d not found", req.positionID)
	}

	err = order.AddPositions(req.positionID, req.quantity, position.Price)
	if err != nil {

		return fmt.Errorf("add position to order with id %d failed", req.orderID)
	}

	_, err = order.Value()
	if err != nil {

		return fmt.Errorf("failed to serialized positions %d", err)
	}

	return nil
}
