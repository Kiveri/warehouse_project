package order_db

import "warehouse_project/internal/domain/model"

type OrderRepo struct {
	orders map[int64]*model.Order
	nextID int64
}

func (or *OrderRepo) getNextID() int64 {
	nextID := or.nextID
	or.nextID++

	return nextID
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{
		orders: make(map[int64]*model.Order),
		nextID: 1,
	}
}
