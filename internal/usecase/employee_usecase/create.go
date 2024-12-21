package employee_usecase

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

type CreateEmployeeReq struct {
	Name  string
	Phone string
	Email string
	Role  model.EmployeeRole
}

func (eu *EmployeeUseCase) CreateEmployee(req CreateEmployeeReq) (*model.Employee, error) {
	now := eu.timer.Now()
	employee := model.NewEmployee(req.Name, req.Phone, req.Email, req.Role, now)
	employee, err := eu.employeeRepo.CreateEmployee(employee)
	if err != nil {
		return nil, fmt.Errorf("employeeRepo.CreateEmployee: %w", err)
	}

	return employee, nil
}
