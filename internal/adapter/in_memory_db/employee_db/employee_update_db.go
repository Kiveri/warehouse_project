package employee_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (r *EmployeeRepo) UpdateEmployee(employee *model.Employee) error {
	if _, exists := r.employees[employee.ID]; !exists {
		return fmt.Errorf("employee with id %s does not exist", employee.ID)
	}
	r.employees[employee.ID] = employee
	return nil
}
