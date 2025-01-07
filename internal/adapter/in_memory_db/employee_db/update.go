package employee_db

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (er *EmployeeRepo) UpdateEmployee(employee *model.Employee) (*model.Employee, error) {
	if _, exists := er.employeesMap[employee.ID]; !exists {
		return nil, fmt.Errorf("employee with id %v does not exist", employee.ID)
	}
	er.employeesMap[employee.ID] = employee

	return employee, nil
}
