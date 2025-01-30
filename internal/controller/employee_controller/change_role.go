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

func (c *Controller) ChangeRole(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	employeeID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("employee id not present", "id"))

		return
	}

	decoder := json.NewDecoder(r.Body)
	var req changeRoleRequest
	err = decoder.Decode(&req)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("employee not found", "id"))

		return
	}

	if validationError := validateChangeRoleRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	updateEmployeeRole, err := c.employeeUseCase.UpdateEmployee(employee_usecase.UpdateEmployeeReq{
		ID:   employeeID,
		Role: req.Role,
	})
	if err != nil {
		if errors.Is(err, employees.NotFound) {
			controller.ValidationErrorRespond(w, controller.NewValidationError("employee not found", "id"))

			return
		}

		controller.InternalServer(w, err)

		return
	}

	controller.Validation(w, http.StatusOK, updateEmployeeRole)
}

func validateChangeRoleRequest(req changeRoleRequest) *controller.ValidationError {
	if req.Role <= 0 || req.Role >= 4 {
		return controller.NewValidationError("role 1 - specialist, 2 - leader, 3 - manager", "role")
	}

	return nil
}
