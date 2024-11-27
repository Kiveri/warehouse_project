package employee_usecase

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

type UpdateEmployeeRequest struct {
	ID   int64
	Role model.EmployeeRole
}

func (eu *EmployeeUseCase) UpdateEmployee(req UpdateEmployeeRequest) error {
	employee, err := eu.employeeRepo.FindEmployee(req.ID)
	if err != nil {
		return fmt.Errorf("employeeRepo.FindEmployee: %w", err)
	}

	employee.ChangeRole(req.Role)

	if err = eu.employeeRepo.UpdateEmployee(employee); err != nil {
		return fmt.Errorf("employeeRepo.UpdateEmployee: %w", err)
	}

	return nil
}
