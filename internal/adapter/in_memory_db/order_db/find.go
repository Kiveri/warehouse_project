package order_db

import "warehouse_project/internal/domain/model"

func (or *OrderRepo) FindOrder(id int64) (*model.Order, bool) {
	order, exists := or.orders[id]
	return order, exists
}
