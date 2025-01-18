package orders

import (
	"context"
	"fmt"
)

func (r *Repo) DeleteOrder(id int64) error {
	order, err := r.cluster.Conn.Exec(context.Background(),
		"DELETE FROM orders WHERE id = $1 ", id)

	if err != nil {
		return fmt.Errorf("r.cluster.QueryRow: %w", err)
	}

	if order.RowsAffected() == 0 {
		return fmt.Errorf("position with id %d not found", id)
	}

	return nil
}
