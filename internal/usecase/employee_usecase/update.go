package employee_usecase

import "warehouse_project/internal/domain/model"

func (eu *EmployeeUseCase) UpdateEmployee(employee *model.Employee) error {
	return eu.er.UpdateEmployee(employee)
}
