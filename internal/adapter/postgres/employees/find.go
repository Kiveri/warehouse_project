package employees

import (
	"context"
	"errors"
	"fmt"

	"warehouse_project/internal/domain/model"

	"github.com/jackc/pgx/v5"
)

func (r *Repo) FindEmployee(id int) (*model.Employee, error) {
	var employee model.Employee

	query := "SELECT id, name, phone, email, role, created_at, updated_at " +
		"FROM employees " +
		"WHERE id = $1"

	err := r.cluster.Conn.QueryRow(context.Background(), query, id).
		Scan(
			&employee.ID,
			&employee.Name,
			&employee.Phone,
			&employee.Email,
			&employee.Role,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no client found with id %d", id)
		}
		return nil, fmt.Errorf("FindClient: %w", err)
	}

	return &employee, nil
}
