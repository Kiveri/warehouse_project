package employee_usecase

func (eu *EmployeeUseCase) DeleteEmployeeUC(id int) error {
	return eu.er.DeleteEmployee(id)
}
