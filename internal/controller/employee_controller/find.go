package employee_controller

import (
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/employees"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/employee_usecase"
)

func (c *Controller) Find(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	employeeID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("employee id not present", "id"))

		return
	}

	findEmployee, err := c.employeeUseCase.FindEmployee(employee_usecase.FindEmployeeReq{
		ID: employeeID,
	})
	if err != nil {
		if errors.Is(err, employees.NotFound) {
			controller.ValidationErrorRespond(w, controller.NewValidationError("employee not found", "id"))

			return
		}

		controller.InternalServer(w, err)

		return
	}

	controller.Validation(w, http.StatusOK, findEmployee)
}
