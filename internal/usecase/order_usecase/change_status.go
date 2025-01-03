package order_usecase

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

type UpdateOrderReq struct {
	ID     int64
	Status model.OrderStatus
}

func (ou *OrderUseCase) UpdateOrder(req UpdateOrderReq) error {
	order, err := ou.orderRepo.FindOrder(req.ID)
	if err != nil {
		return fmt.Errorf("orderRepo.FindOrder: %w", err)
	}

	order.ChangeStatus(req.Status)

	if err = ou.orderRepo.UpdateOrder(order); err != nil {
		return fmt.Errorf("orderRepo.UpdateOrder: %w", err)
	}

	return nil
}
