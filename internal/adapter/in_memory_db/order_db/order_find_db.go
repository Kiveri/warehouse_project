package order_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (r *OrderRepo) FindOrder(id int) (*model.Order, error) {
	order, exists := r.orders[id]
	if !exists {
		return nil, fmt.Errorf("order not found")
	}
	return order, nil
}
