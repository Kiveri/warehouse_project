package clients

import (
	"context"
	"fmt"
)

func (r *Repo) DeleteClient(id int64) error {

	client, err := r.cluster.Conn.Exec(context.Background(),
		"DELETE FROM clients WHERE id = $1", id)

	if err != nil {
		return fmt.Errorf("r.cluster.QueryRow: %w", err)
	}

	if client.RowsAffected() == 0 {
		return fmt.Errorf("client with id %d not found", id)
	}

	return nil
}
