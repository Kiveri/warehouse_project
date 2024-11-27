package order_db

import "warehouse_project/internal/domain/model"

type OrderRepo struct {
	orders map[int]model.Order
	nextID int
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{
		orders: make(map[int]model.Order),
		nextID: 1,
	}
}
