package order_usecase

/*import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

type UpdateOrderReq struct {
	ID     int64
	Status model.OrderStatus
}

func (ou *OrderUseCase) UpdateOrder(req UpdateOrderReq) (*model.Order, error) {
	order, err := ou.orderRepo.FindOrder(req.ID)
	if err != nil {
		return nil, fmt.Errorf("orderRepo.FindOrder: %w", err)
	}

	order.ChangeStatus(req.Status, ou.timer.Now())

	if _, err = ou.orderRepo.UpdateOrder(order); err != nil {
		return nil, fmt.Errorf("orderRepo.UpdateOrder: %w", err)
	}

	return order, nil
}

*/
