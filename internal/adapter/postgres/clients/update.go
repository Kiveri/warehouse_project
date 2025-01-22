package clients

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) UpdateClient(client *model.Client) (*model.Client, error) {
	query := `
		UPDATE clients SET name = $1, phone = $2, email = $3, home_address = $4, updated_at = $5 
		WHERE id = $6 
		RETURNING id, name, phone, email, home_address, created_at, updated_at
		`

	err := r.cluster.Conn.QueryRow(context.Background(), query,
		client.Name,
		client.Phone,
		client.Email,
		client.HomeAddress,
		client.UpdatedAt,
		client.ID,
	)

	if err != nil {

		return nil, fmt.Errorf("UpdateClient: failed to update client: %v", err)
	}

	return client, nil
}
