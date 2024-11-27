package employee_usecase

func (eu *EmployeeUseCase) DeleteEmployee(id int) error {
	return eu.er.DeleteEmployee(id)
}
