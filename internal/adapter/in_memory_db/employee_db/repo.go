package employee_db

import (
	"warehouse_project/internal/domain/model"
)

type EmployeeRepo struct {
	employees map[int64]*model.Employee
	nextID    int64
}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{
		employees: make(map[int64]*model.Employee),
		nextID:    1,
	}
}
