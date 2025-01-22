package employees

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) UpdateEmployee(employee *model.Employee) (*model.Employee, error) {
	query := `
		UPDATE employees SET name = $1, phone = $2, email = $3, role = $4, updated_at = $5 
		WHERE id = $6 
		RETURNING id, name, phone, email, role, created_at, updated_at
		`

	err := r.cluster.Conn.QueryRow(context.Background(), query,
		employee.Name,
		employee.Phone,
		employee.Email,
		employee.Role,
		employee.UpdatedAt,
		employee.ID,
	)

	if err != nil {

		return nil, fmt.Errorf("UpdateEmployee: failed to update employee: %v", err)
	}

	return employee, nil
}
