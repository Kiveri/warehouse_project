package employee_controller

import (
	"encoding/json"
	"net/http"

	"warehouse_project/internal/controller"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/employee_usecase"
)

type createEmployeeRequest struct {
	Name  string             `json:"name"`
	Phone string             `json:"phone"`
	Email string             `json:"email"`
	Role  model.EmployeeRole `json:"role"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req createEmployeeRequest
	err := decoder.Decode(&req)
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}

	if validationError := validateCreateEmployeeRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	employee, err := c.employeeUseCase.CreateEmployee(employee_usecase.CreateEmployeeReq{
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
		Role:  req.Role,
	})
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}

	controller.Respond(w, http.StatusOK, employee)
}

func validateCreateEmployeeRequest(req createEmployeeRequest) *controller.ValidationError {
	if req.Name == "" {
		return controller.NewValidationError("name is required", "name")
	}
	if req.Role <= 0 || req.Role >= 4 {
		return controller.NewValidationError("role 1 - specialist, 2 - leader, 3 - manager", "role")
	}

	return nil
}
