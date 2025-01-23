package orders

import (
	"context"
	"fmt"
)

func (r *Repo) DeleteOrder(id int64) error {
	_, err := r.cluster.Conn.Exec(context.Background(),
		"DELETE FROM orders WHERE id = $1 ", id)

	if err != nil {
		return fmt.Errorf("r.cluster.QueryRow: %w", err)
	}

	return nil
}
