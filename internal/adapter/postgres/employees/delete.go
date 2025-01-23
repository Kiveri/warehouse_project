package employees

import (
	"context"
	"fmt"
)

func (r *Repo) DeleteEmployee(id int64) error {
	_, err := r.cluster.Conn.Exec(context.Background(),
		"DELETE FROM employees WHERE id = $1 ", id)

	if err != nil {
		return fmt.Errorf("r.cluster.QueryRow: %w", err)
	}

	return nil
}
