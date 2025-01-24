package employee_controller

import (
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/employee_usecase"
)

type (
	Controller struct {
		employeeUseCase employeeUseCase
	}

	employeeUseCase interface {
		CreateEmployee(req employee_usecase.CreateEmployeeReq) (*model.Employee, error)
		FindEmployee(req employee_usecase.FindEmployeeReq) (*model.Employee, error)
		UpdateEmployee(req employee_usecase.UpdateEmployeeReq) (*model.Employee, error)
		DeleteEmployee(req employee_usecase.DeleteEmployeeReq) error
	}
)

func NewController(employeeUseCase employeeUseCase) *Controller {
	return &Controller{
		employeeUseCase: employeeUseCase,
	}
}
