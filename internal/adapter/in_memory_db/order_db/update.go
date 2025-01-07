package order_db

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (or *OrderRepo) UpdateOrder(order *model.Order) (*model.Order, error) {
	if _, exists := or.ordersMap[order.ID]; exists {
		return nil, fmt.Errorf("order with id %v does not exists", order.ID)
	}
	or.ordersMap[order.ID] = order

	return order, nil
}
