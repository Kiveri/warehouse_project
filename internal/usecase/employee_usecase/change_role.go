package employee_usecase

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

type UpdateEmployeeReq struct {
	ID   int64
	Role model.EmployeeRole
}

func (eu *EmployeeUseCase) UpdateEmployee(req UpdateEmployeeReq) (*model.Employee, error) {
	employee, err := eu.employeeRepo.FindEmployee(req.ID)
	if err != nil {
		return nil, fmt.Errorf("employeeRepo.FindEmployee: %w", err)
	}

	employee.ChangeRole(req.Role, eu.timer.Now())

	if _, err = eu.employeeRepo.UpdateEmployee(employee); err != nil {
		return nil, fmt.Errorf("employeeRepo.UpdateEmployee: %w", err)
	}

	return employee, nil
}
