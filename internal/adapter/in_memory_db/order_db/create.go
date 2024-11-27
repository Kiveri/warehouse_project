package order_db

import "warehouse_project/internal/domain/model"

func (or *OrderRepo) CreateOrder(order *model.Order) *model.Order {
	order.ID = or.nextID
	or.orders[order.ID] = order
	or.nextID++
	return order
}
