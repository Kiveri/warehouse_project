package positions

import (
	"context"
	"errors"
	"fmt"

	"warehouse_project/internal/domain/model"

	"github.com/jackc/pgx/v5"
)

func (r *Repo) FindPosition(id int64) (*model.Position, error) {
	var position model.Position

	query := `
		SELECT id, name, barcode, price, position_type, created_at, updated_at 
		FROM positions 
		WHERE id = $1
		`

	err := r.cluster.Conn.QueryRow(context.Background(), query, id).
		Scan(
			&position.ID,
			&position.Name,
			&position.Barcode,
			&position.Price,
			&position.PositionType,
			&position.CreatedAt,
			&position.UpdatedAt,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no position found with id %d", id)
		}
		return nil, fmt.Errorf("FindPosition: %w", err)
	}

	return &position, nil
}
