package employee_db

import (
	"fmt"
)

func (er *EmployeeRepo) DeleteEmployee(id int64) error {
	if _, exists := er.employeesMap[id]; !exists {
		return fmt.Errorf("employee not found")
	}
	delete(er.employeesMap, id)

	return nil
}
