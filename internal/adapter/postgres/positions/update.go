package positions

import (
	"context"
	"fmt"
	"time"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) UpdatePosition(position *model.Position) (*model.Position, error) {
	now := time.Now()
	var updatedPosition model.Position

	query := `
		UPDATE positions SET name = $1, barcode = $2, price = $3, position_type = $4, updated_at = $5 
		WHERE id = $6 
		RETURNING id, name, barcode, price, position_type, created_at, updated_at
		`

	if position.UpdatedAt.IsZero() {
		position.UpdatedAt = now
	}

	err := r.cluster.Conn.QueryRow(context.Background(), query,
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
