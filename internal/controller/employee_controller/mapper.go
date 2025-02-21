package employee_controller

import (
	"time"

	"warehouse_project/internal/domain/model"
)

type employeeResponse struct {
	ID        int64              `json:"id"`
	Name      string             `json:"name"`
	Phone     string             `json:"phone"`
	Email     string             `json:"email"`
	Role      model.EmployeeRole `json:"role"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func mapEmployeeToResponse(employee *model.Employee) *employeeResponse {
	return &employeeResponse{
		ID:        employee.ID,
		Name:      employee.Name,
		Phone:     employee.Phone,
		Email:     employee.Email,
		Role:      employee.Role,
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
	}
}
