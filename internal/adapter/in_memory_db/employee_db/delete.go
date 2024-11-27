package employee_db

import (
	"fmt"
)

func (er *EmployeeRepo) DeleteEmployee(id int) error {
	if _, exists := er.employees[id]; !exists {
		return fmt.Errorf("Employee not found")
	}
	delete(er.employees, id)
	return nil
}
