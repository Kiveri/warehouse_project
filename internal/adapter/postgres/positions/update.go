package positions

import (
	"context"
	"fmt"
	"time"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) UpdatePosition(position *model.Position) (*model.Position, error) {
	var updatedPosition model.Position

	query := "UPDATE positions SET name = $1, phone = $2, email = $3, home_address = $4, updated_at = $5 " +
		"WHERE id = $6 " +
		"RETURNING id, name, phone, email, role, created_at, updated_at"

	if position.UpdatedAt.IsZero() {
		position.UpdatedAt = time.Now().UTC()
	}

	err := r.cluster.Conn.QueryRow(context.Background(),
		query,
		position.Name,
		position.Barcode,
		position.Price,
		position.PositionType,
		position.UpdatedAt,
		position.ID,
	).Scan(
		&updatedPosition.ID,
		&updatedPosition.Name,
		&updatedPosition.Barcode,
		&updatedPosition.Price,
		&updatedPosition.PositionType,
		&updatedPosition.CreatedAt,
		&updatedPosition.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("updatedPosition: failed to update position: %w", err)
	}

	return &updatedPosition, nil
}
