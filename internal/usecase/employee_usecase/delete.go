package employee_usecase

import "fmt"

func (eu *EmployeeUseCase) DeleteEmployee(id int64) error {
	if err := eu.employeeRepo.DeleteEmployee(id); err != nil {
		return fmt.Errorf("employeeRepo.DeleteEmployee: %w", err)
	}

	return nil
}
