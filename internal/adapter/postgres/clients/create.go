package clients

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) CreateClient(c *model.Client) (*model.Client, error) {
	if err := r.cluster.Conn.QueryRow(context.Background(),
		"INSERT INTO clients (name, phone, email, home_address, created_at, updated_at) "+
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		c.Name, c.Phone, c.Email, c.HomeAddress, c.CreatedAt, c.UpdatedAt).
		Scan(&c.ID); err != nil {
		return nil, fmt.Errorf("Conn.QueryRow: %w", err)
	}

	return c, nil
}
