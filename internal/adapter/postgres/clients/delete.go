package clients

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) DeleteClient(id int64) (*model.Client, error) {
	var deletedClient model.Client

	err := r.cluster.Conn.QueryRow(context.Background(),
		"DELETE FROM clients WHERE id = $1 "+
			"RETURNING id, name, phone, email, home_address, created_at, updated_at", id).
		Scan(
			&deletedClient.ID,
			&deletedClient.Name,
			&deletedClient.Phone,
			&deletedClient.Email,
			&deletedClient.HomeAddress,
			&deletedClient.CreatedAt,
			&deletedClient.UpdatedAt,
		)

	if err != nil {
		return nil, fmt.Errorf("r.cluster.QueryRow: %w", err)
	}

	return &deletedClient, nil
}
