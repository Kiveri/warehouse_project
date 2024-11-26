package employee_usecase

func (u *EmployeeUseCase) DeleteEmployeeUC(id int) error {
	return u.r.DeleteEmployee(id)
}
