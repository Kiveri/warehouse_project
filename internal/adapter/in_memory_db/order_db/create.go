package order_db

import "warehouse_project/internal/domain/model"

func (or *OrderRepo) CreateOrder(order *model.Order) (*model.Order, error) {
	order.ID = or.getNextID()
	or.orders[order.ID] = order

	return order, nil
}
