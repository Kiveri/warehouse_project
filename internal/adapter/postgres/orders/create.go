package orders

import (
	"context"
	"encoding/json"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) CreateOrder(order *model.Order) (*model.Order, error) {
	query := `
			INSERT INTO orders (positions, employee_id, client_id, status, delivery_type, total, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
		`

	positionsJSON, err := json.Marshal(order.Positions)
	if err != nil {
		return nil, fmt.Errorf("error encoding positions: %w", err)
	}

	err = r.cluster.Conn.QueryRow(context.Background(), query,
		positionsJSON,
		order.EmployeeID,
		order.ClientID,
		order.Status,
		order.DeliveryType,
		order.Total,
		order.CreatedAt,
		order.UpdatedAt,
	).Scan(&order.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	return order, nil
}
