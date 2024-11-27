package order_usecase

import (
	"warehouse_project/internal/adapter/in_memory_db/order_db"
	"warehouse_project/internal/adapter/in_memory_db/position_db"
)

type OrderUseCase struct {
	positionRepo *position_db.PositionRepo
	orderRepo    *order_db.OrderRepo
}

func NewOrderUseCase(or *order_db.OrderRepo, pr *position_db.PositionRepo) *OrderUseCase {
	return &OrderUseCase{
		positionRepo: pr,
		orderRepo:    or,
	}
}
