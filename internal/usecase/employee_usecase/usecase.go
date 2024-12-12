package employee_usecase

import (
	"warehouse_project/internal/domain/model"
)

type employeeRepo interface {
	CreateEmployee(employee *model.Employee) (*model.Employee, error)
	DeleteEmployee(id int64) error
	FindEmployee(id int64) (*model.Employee, error)
	UpdateEmployee(employee *model.Employee) error
}

type EmployeeUseCase struct {
	employeeRepo employeeRepo
}

func NewEmployeeUseCase(employeeRepo employeeRepo) *EmployeeUseCase {
	return &EmployeeUseCase{
		employeeRepo: employeeRepo,
	}
}
