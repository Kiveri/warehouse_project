package employee_usecase

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

type FindEmployeeReq struct {
	ID int64
}

func (eu *EmployeeUseCase) FindEmployee(req FindEmployeeReq) (*model.Employee, error) {
	employee, err := eu.employeeRepo.FindEmployee(req.ID)
	if err != nil {
		return nil, fmt.Errorf("employeeRepo.FindEmployee: %w", err)
	}

	return employee, nil
}
