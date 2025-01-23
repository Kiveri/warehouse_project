package orders

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) CreateOrder(order *model.Order) (*model.Order, error) {
	query := `
			INSERT INTO orders (positions, employee_id, client_id, status, delivery_type, total, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
		`

	err := r.cluster.Conn.QueryRow(context.Background(), query,
		order.Positions,
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
