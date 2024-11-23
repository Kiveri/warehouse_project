package employee_db

import (
	"fmt"
)

func (r *EmployeeRepo) DeleteEmployee(id int) error {
	if _, exists := r.employees[id]; !exists {
		return fmt.Errorf("Employee not found")
	}
	delete(r.employees, id)
	return nil
}
