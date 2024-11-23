package employee_usecase

import "warehouse_project/internal/domain/model"

func (u *EmployeeUseCase) FindEmployeeUC(id int) (*model.Employee, error) {
	return u.r.FindEmployee(id)
}
