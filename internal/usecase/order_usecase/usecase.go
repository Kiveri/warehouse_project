package order_usecase

import (
	"warehouse_project/internal/adapter/in_memory_db/employee_db"
	"warehouse_project/internal/adapter/in_memory_db/order_db"
	"warehouse_project/internal/adapter/in_memory_db/position_db"
)

type OrderUseCase struct {
	positionRepo *position_db.PositionRepo
	orderRepo    *order_db.OrderRepo
	employeeRepo *employee_db.EmployeeRepo
}

func NewOrderUseCase(orderRepo *order_db.OrderRepo, positionRepo *position_db.PositionRepo, employeeRepo *employee_db.EmployeeRepo) *OrderUseCase {
	return &OrderUseCase{
		positionRepo: positionRepo,
		orderRepo:    orderRepo,
		employeeRepo: employeeRepo,
	}
}
