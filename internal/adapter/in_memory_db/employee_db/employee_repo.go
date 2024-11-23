package employee_db

import (
	"warehouse_project/internal/domain/model"
)

type EmployeeRepo struct {
	employees map[int]*model.Employee
	nextID    int
}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{
		employees: make(map[int]*model.Employee),
		nextID:    1,
	}
}
