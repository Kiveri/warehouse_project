package clients

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) CreateClient(client *model.Client) (*model.Client, error) {
	err := r.cluster.Conn.QueryRow(context.Background(),
		"INSERT INTO clients (name, phone, email, home_address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) "+
			"RETURNING id",
		client.Name, client.Phone, client.Email, client.HomeAddress, client.CreatedAt, client.UpdatedAt).
		Scan(&client.ID)

	if err != nil {

		return nil, fmt.Errorf("Conn.QueryRow: %w", err)
	}

	return client, nil
}
