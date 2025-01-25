package employee_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/employees"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/employee_usecase"
)

type changeRoleRequest struct {
	Role model.EmployeeRole `json:"role"`
}

func (c *Controller) ChangeRole() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		employeeID, err := strconv.ParseInt(id, 0, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(controller.NewValidationError("employee id not present", "id"))

			return
		}

		decoder := json.NewDecoder(r.Body)
		var req changeRoleRequest
		err = decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		if validationError := validateChangeRoleRequest(req); validationError != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(validationError)

			return
		}

		updateEmployeeRole, err := c.employeeUseCase.UpdateEmployee(employee_usecase.UpdateEmployeeReq{
			ID:   employeeID,
			Role: req.Role,
		})
		if err != nil {
			if errors.Is(err, employees.NotFound) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(controller.NewNotFoundError("employee id not found"))

				return
			}

			http.Error(w, err.Error(), 500)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updateEmployeeRole)

		return
	}
}

func validateChangeRoleRequest(req changeRoleRequest) *controller.ValidationError {
	if req.Role <= 0 || req.Role >= 4 {
		return controller.NewValidationError("role 1 - specialist, 2 - leader, 3 - manager", "role")
	}

	return nil
}
