package clients

import (
	"context"
	"errors"
	"fmt"

	"warehouse_project/internal/domain/model"

	"github.com/jackc/pgx/v5"
)

func (r *Repo) FindClient(id int64) (*model.Client, error) {
	var client model.Client

	query := `
		SELECT id, name, phone, email, home_address, created_at, updated_at 
		FROM clients WHERE id = $1
		`

	err := r.cluster.Conn.QueryRow(context.Background(), query, id).
		Scan(
			&client.ID,
			&client.Name,
			&client.Phone,
			&client.Email,
			&client.HomeAddress,
			&client.CreatedAt,
			&client.UpdatedAt,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {

			return nil, fmt.Errorf("no client found with id %d", id)
		}

		return nil, fmt.Errorf("FindClient: %w", err)
	}

	return &client, nil
}
