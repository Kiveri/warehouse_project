package employee_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (er *EmployeeRepo) FindEmployee(id int) (*model.Employee, error) {
	employee, exists := er.employees[id]
	if !exists {
		return nil, fmt.Errorf("Employee not found")
	}
	return employee, nil
}
