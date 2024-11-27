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
	Post    model.EmployeeRole
}

func (eu *EmployeeUseCase) CreateEmployee(req CreateEmployeeReq) (*model.Employee, error) {
	now := time.Now()
	employee := model.NewEmployee(req.Name, req.Surname, req.Phone, req.Email, req.Post, now)
	employee, err := eu.er.CreateEmployee(employee)
	if err != nil {
		return nil, fmt.Errorf("CreateEmployeeUC err: %w", err)
	}
	return employee, nil
}
