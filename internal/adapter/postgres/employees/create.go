package employees

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) CreateEmployee(employee *model.Employee) (*model.Employee, error) {
	if err := r.cluster.Conn.QueryRow(context.Background(),
		"INSERT INTO employees (name, phone, email, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		employee.Name, employee.Phone, employee.Email, employee.Role, employee.CreatedAt, employee.UpdatedAt).
		Scan(&employee.ID); err != nil {
		return nil, fmt.Errorf("Conn.QueryRow: %w", err)
	}

	return employee, nil
}
