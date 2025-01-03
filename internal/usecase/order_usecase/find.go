package order_usecase

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (ou *OrderUseCase) FindOrder(id int64) (*model.Order, error) {
	order, err := ou.orderRepo.FindOrder(id)
	if err != nil {
		return nil, fmt.Errorf("orderRepo.FindOrder: %w", err)
	}

	return order, nil
}
