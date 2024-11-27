package employee_usecase

import (
	"fmt"
	"time"
	"warehouse_project/internal/domain/model"
)

type CreateEmployeeReq struct {
	Name    string
	Surname string
	Phone   string
	Email   string
	Role    model.EmployeeRole
}

func (eu *EmployeeUseCase) CreateEmployee(req CreateEmployeeReq) (*model.Employee, error) {
	employee := model.NewEmployee(req.Name, req.Surname, req.Phone, req.Email, req.Role, time.Now())
	employee, err := eu.employeeRepo.CreateEmployee(employee)
	if err != nil {
		return nil, fmt.Errorf("employeeRepo.CreateEmployee: %w", err)
	}

	return employee, nil
}
