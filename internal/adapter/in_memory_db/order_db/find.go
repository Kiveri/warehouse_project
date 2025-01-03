package order_db

import (
	"errors"
	"warehouse_project/internal/domain/model"
)

func (or *OrderRepo) FindOrder(id int64) (*model.Order, error) {
	order, exists := or.ordersMap[id]
	if !exists {
		return nil, errors.New("order not found")
	}

	return order, nil
}
