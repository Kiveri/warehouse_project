package employee_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/employees"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/employee_usecase"
)

func (c *Controller) Find() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		employeeID, err := strconv.ParseInt(id, 0, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(controller.NewValidationError("employee id not present", "id"))

			return
		}

		findEmployee, err := c.employeeUseCase.FindEmployee(employee_usecase.FindEmployeeReq{
			ID: employeeID,
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
		json.NewEncoder(w).Encode(findEmployee)

		return
	}
}
