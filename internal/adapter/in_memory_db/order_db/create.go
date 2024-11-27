package order_db

import "warehouse_project/internal/domain/model"

func (r *OrderRepo) CreateOrder(order *model.Order) (*model.Order, error) {
	order.ID = r.nextID
	r.orders[order.ID] = order
	r.nextID++
	return order, nil
}
