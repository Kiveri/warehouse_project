package employee_usecase

import "warehouse_project/internal/domain/model"

func (eu *EmployeeUseCase) FindEmployee(id int) (*model.Employee, error) {
	return eu.er.FindEmployee(id)
}
