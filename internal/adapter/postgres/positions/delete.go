package positions

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) DeletePosition(id int) (*model.Position, error) {
	var deletedPosition model.Position

	err := r.cluster.Conn.QueryRow(context.Background(),
		"DELETE FROM positions WHERE id = $1 "+
			"RETURNING id, name, barcode, price, position_type, created_at, updated_at", id).
		Scan(
			&deletedPosition.ID,
			&deletedPosition.Name,
			&deletedPosition.Barcode,
			&deletedPosition.Price,
			&deletedPosition.PositionType,
			&deletedPosition.CreatedAt,
			&deletedPosition.UpdatedAt,
		)

	if err != nil {
		return nil, fmt.Errorf("r.cluster.QueryRow: %w", err)
	}

	return &deletedPosition, nil
}
