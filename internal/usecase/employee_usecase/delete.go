package employee_usecase

import "fmt"

type DeleteEmployeeReq struct {
	ID int64
}

func (eu *EmployeeUseCase) DeleteEmployee(req DeleteEmployeeReq) error {
	if err := eu.employeeRepo.DeleteEmployee(req.ID); err != nil {
		return fmt.Errorf("employeeRepo.DeleteEmployee: %w", err)
	}

	return nil
}
