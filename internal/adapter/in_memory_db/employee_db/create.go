package employee_db

import "warehouse_project/internal/domain/model"

func (er *EmployeeRepo) CreateEmployee(employee *model.Employee) (*model.Employee, error) {
	employee.ID = er.nextID
	er.employees[employee.ID] = employee
	er.nextID++

	return employee, nil

}
