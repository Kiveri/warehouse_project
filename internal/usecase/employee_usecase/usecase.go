package employee_usecase

import (
	"time"

	"warehouse_project/internal/domain/model"
)

type EmployeeUseCase struct {
	employeeRepo employeeRepo
	timer        timer
}
type (
	employeeRepo interface {
		CreateEmployee(employee *model.Employee) (*model.Employee, error)
		DeleteEmployee(id int64) error
		FindEmployee(id int64) (*model.Employee, error)
		UpdateEmployee(employee *model.Employee) (*model.Employee, error)
	}
	timer interface {
		Now() time.Time
	}
)

func NewEmployeeUseCase(employeeRepo employeeRepo, timer timer) *EmployeeUseCase {
	return &EmployeeUseCase{
		employeeRepo: employeeRepo,
		timer:        timer,
	}
}
