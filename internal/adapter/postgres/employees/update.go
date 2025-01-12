package employees

import (
	"context"
	"fmt"
	"time"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) UpdateEmployee(employee *model.Employee) (*model.Employee, error) {
	var updatedEmployee model.Employee

	query := "UPDATE employees SET name = $1, phone = $2, email = $3, home_address = $4, updated_at = $5 " +
		"WHERE id = $6 " +
		"RETURNING id, name, phone, email, role, created_at, updated_at"

	// Установка метки времени обновления, если она еще не установлена
	if employee.UpdatedAt.IsZero() {
		employee.UpdatedAt = time.Now().UTC()
	}

	err := r.cluster.Conn.QueryRow(context.Background(),
		query,
		employee.Name,
		employee.Phone,
		employee.Email,
		employee.Role,
		employee.UpdatedAt,
		employee.ID,
	).Scan(
		&updatedEmployee.ID,
		&updatedEmployee.Name,
		&updatedEmployee.Phone,
		&updatedEmployee.Email,
		&updatedEmployee.Role,
		&updatedEmployee.CreatedAt,
		&updatedEmployee.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("UpdateEmployee: failed to update employee: %w", err)
	}

	return &updatedEmployee, nil
}
