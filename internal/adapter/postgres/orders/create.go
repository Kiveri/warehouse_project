package orders

import (
	"context"
	"encoding/json"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	query := `
			INSERT INTO orders (positions, employeeID, clientID, status, deliveryType, total, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
		`

	positionsJSON, err := json.Marshal(order.Positions)
	if err != nil {
		return nil, fmt.Errorf("error encoding positions: %w", err)
	}

	err = r.cluster.Conn.QueryRow(ctx, query,
		positionsJSON,
		order.EmployeeID,
		order.ClientID,
		order.Status,
		order.Total,
		order.CreatedAt,
		order.UpdatedAt,
	).Scan(&order.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	return order, nil
}
