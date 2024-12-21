package order_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (or *OrderRepo) FindAllByIDs(id []int64) ([]*model.Order, error) {
	orders := make([]*model.Order, 0, len(id))

	for _, ordersId := range id {
		if order, exists := or.ordersMap[ordersId]; exists {
			orders = append(orders, order)
		} else {
			return nil, fmt.Errorf("order with id %d not found", ordersId)
		}
	}

	return orders, nil
}
