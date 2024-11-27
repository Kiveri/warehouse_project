package employee_usecase

import "warehouse_project/internal/adapter/in_memory_db/employee_db"

type EmployeeUseCase struct {
	employeeRepo *employee_db.EmployeeRepo
}

func NewEmployeeUseCase(employeeRepo *employee_db.EmployeeRepo) *EmployeeUseCase {
	return &EmployeeUseCase{
		employeeRepo: employeeRepo,
	}
}
