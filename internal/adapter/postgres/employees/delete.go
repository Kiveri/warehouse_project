package employees

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) DeleteEmployee(id int64) (*model.Employee, error) {
	var deletedEmployee model.Employee

	err := r.cluster.Conn.QueryRow(context.Background(),
		"DELETE FROM employees WHERE id = $1 "+
			"RETURNING id, name, phone, email, role, created_at, updated_at", id).
		Scan(
			&deletedEmployee.ID,
			&deletedEmployee.Name,
			&deletedEmployee.Phone,
			&deletedEmployee.Email,
			&deletedEmployee.Role,
			&deletedEmployee.CreatedAt,
			&deletedEmployee.UpdatedAt,
		)

	if err != nil {
		return nil, fmt.Errorf("r.cluster.QueryRow: %w", err)
	}

	return &deletedEmployee, nil
}
