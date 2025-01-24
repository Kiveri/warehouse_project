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

func (c *Controller) Create() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req createEmployeeRequest
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		if validationError := validateCreateEmployeeRequest(req); validationError != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(validationError)

			return
		}

		employee, err := c.employeeUseCase.CreateEmployee(employee_usecase.CreateEmployeeReq{
			Name:  req.Name,
			Phone: req.Phone,
			Email: req.Email,
			Role:  req.Role,
		})
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)

		return
	}
}

func validateCreateEmployeeRequest(req createEmployeeRequest) *controller.ValidationError {
	if req.Name == "" {
		return controller.NewValidationError("name is required", "name")
	}

	return nil
}
