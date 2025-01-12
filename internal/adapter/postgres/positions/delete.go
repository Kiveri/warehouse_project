package positions

import (
	"context"
	"fmt"
)

func (r *Repo) DeletePosition(id int64) error {
	position, err := r.cluster.Conn.Exec(context.Background(),
		"DELETE FROM positions WHERE id = $1 ", id)

	if err != nil {
		return fmt.Errorf("r.cluster.QueryRow: %w", err)
	}

	if position.RowsAffected() == 0 {
		return fmt.Errorf("position with id %d not found", id)
	}

	return nil
}
