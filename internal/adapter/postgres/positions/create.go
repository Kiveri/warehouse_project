package positions

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) CreateClient(position *model.Position) (*model.Position, error) {
	if err := r.cluster.Conn.QueryRow(context.Background(),
		"INSERT INTO positions (name, barcode, price, position_type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) "+
			"RETURNING id",
		position.Name, position.Barcode, position.Price, position.PositionType, position.CreatedAt, position.UpdatedAt).
		Scan(&position.ID); err != nil {
		return nil, fmt.Errorf("Conn.QueryRow: %w", err)

	}

	return position, nil
}
