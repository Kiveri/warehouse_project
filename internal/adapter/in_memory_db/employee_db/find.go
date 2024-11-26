package employee_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (r *EmployeeRepo) FindEmployee(id int) (*model.Employee, error) {
	employee, exists := r.employees[id]
	if !exists {
		return nil, fmt.Errorf("Employee not found")
	}
	return employee, nil
}
