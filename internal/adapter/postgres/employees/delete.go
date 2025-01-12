package employees

import (
	"context"
	"fmt"
)

func (r *Repo) DeleteEmployee(id int64) error {
	employee, err := r.cluster.Conn.Exec(context.Background(),
		"DELETE FROM employees WHERE id = $1 ", id)

	if err != nil {
		return fmt.Errorf("r.cluster.QueryRow: %w", err)
	}

	if employee.RowsAffected() == 0 {
		return fmt.Errorf("employee with id %d not found", id)
	}

	return nil
}
