package employee_controller

import (
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/employees"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/employee_usecase"
)

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	employeeID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("employee id not present", "id"))

		return
	}

	err = c.employeeUseCase.DeleteEmployee(employee_usecase.DeleteEmployeeReq{
		ID: employeeID,
	})
	if err != nil {
		if errors.Is(err, employees.NotFound) {
			controller.NotFoundErrorRespond(w, controller.NewNotFoundError("employee not found"))

			return
		}

		controller.InternalServerErrorRespond(w, err)

	}
}
