package employee_db

import "warehouse_project/internal/domain/model"

func (er *EmployeeRepo) CreateEmployee(employee *model.Employee) (*model.Employee, error) {
	employee.ID = er.getNextID()
	er.employeesMap[employee.ID] = employee

	return employee, nil
}
