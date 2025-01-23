package positions

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) UpdatePosition(position *model.Position) (*model.Position, error) {

	query := `
		UPDATE positions SET name = $1, barcode = $2, price = $3, position_type = $4, updated_at = $5 
		WHERE id = $6 
		RETURNING id, name, barcode, price, position_type, created_at, updated_at
		`

	err := r.cluster.Conn.QueryRow(context.Background(), query,
		position.Name,
		position.Barcode,
		position.Price,
		position.PositionType,
		position.UpdatedAt,
		position.ID,
	)

	if err != nil {

		return nil, fmt.Errorf("updatedPosition: failed to update position: %w", err)
	}

	return position, nil
}
