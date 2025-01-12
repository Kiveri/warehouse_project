package orders

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) CreateOrder(ctx context.Context, o *model.Order) error {
	query := `
			INSERT INTO orders (positions, employeeID, clientID, status, deliveryType, total)
		VALUES ($1, $2, $3, $4, $5, $6)
		`

	_, err := r.cluster.Conn.Exec(ctx, query, o.ID, o.Positions)
	if err != nil {
		return fmt.Errorf("failed to save order: %w", err)
	}

	return nil
}
