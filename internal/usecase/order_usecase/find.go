package order_usecase

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

type FindOrderReq struct {
	ID int64
}

func (ou *OrderUseCase) FindOrder(req FindOrderReq) (*model.Order, error) {
	order, err := ou.orderRepo.FindOrder(req.ID)
	if err != nil {
		return nil, fmt.Errorf("orderRepo.FindOrder: %w", err)
	}

	return order, nil
}
