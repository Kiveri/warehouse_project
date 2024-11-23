package employee_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (r *EmployeeRepo) FindEmployee(id int) (*model.Employee, error) {
	eployee, exists := r.employees[id]
	if !exists {
		return nil, fmt.Errorf("Employee not found")
	}
	return eployee, nil
}
