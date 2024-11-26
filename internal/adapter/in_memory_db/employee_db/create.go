package employee_db

import "warehouse_project/internal/domain/model"

func (r *EmployeeRepo) CreateEmployee(employee *model.Employee) (*model.Employee, error) {
	employee.ID = r.nextID
	r.employees[employee.ID] = employee
	r.nextID++

	return employee, nil

}
