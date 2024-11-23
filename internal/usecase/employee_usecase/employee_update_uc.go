package employee_usecase

import "warehouse_project/internal/domain/model"

func (u *EmployeeUseCase) UpdateEmployeeUC(employee *model.Employee) error {
	return u.r.UpdateEmployee(employee)
}
