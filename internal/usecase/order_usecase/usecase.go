package order_usecase

import "warehouse_project/internal/adapter/in_memory_db/order_db"

type OrderUseCase struct {
	r *order_db.OrderRepo
}

func NewOrderUseCase(r *order_db.OrderRepo) *OrderUseCase {
	return &OrderUseCase{r}
}
