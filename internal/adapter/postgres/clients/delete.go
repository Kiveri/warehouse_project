package clients

import (
	"context"
	"fmt"
)

func (r *Repo) DeleteClient(id int64) error {
	_, err := r.cluster.Conn.Exec(context.Background(),
		"DELETE FROM clients WHERE id = $1", id)

	if err != nil {
		return fmt.Errorf("r.cluster.QueryRow: %w", err)
	}

	return nil
}
