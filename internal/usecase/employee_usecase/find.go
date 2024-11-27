package employee_usecase

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (eu *EmployeeUseCase) FindEmployee(id int64) (*model.Employee, error) {
	employee, err := eu.employeeRepo.FindEmployee(id)
	if err != nil {
		return nil, fmt.Errorf("employeeRepo.FindEmployee: %w", err)
	}

	return employee, nil
}
