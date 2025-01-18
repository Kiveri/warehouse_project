package orders

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) UpdateOrder(order *model.Order) (*model.Order, error) {
	now := time.Now()
	var updatedOrder model.Order

	query := `
		UPDATE orders SET positions = $1, employee_id = $2, client_id = $3, status = $4, delivery_type = $5,  total = $6, updated_at = $7
        WHERE id = $8
		RETURNING id, positions, employee_id, client_id, status, delivery_type, total, updated_at`

	if order.UpdatedAt.IsZero() {
		order.UpdatedAt = now
	}

	positionsJSON, err := json.Marshal(order.Positions)
	if err != nil {
		return nil, fmt.Errorf("error marshalling positions: %w", err)
	}

	err = r.cluster.Conn.QueryRow(context.Background(), query,
		positionsJSON,
		order.EmployeeID,
		order.ClientID,
		order.Status,
		order.DeliveryType,
		order.Total,
		order.UpdatedAt,
		order.ID,
	).Scan(
		&updatedOrder.ID,
		&positionsJSON,
		&updatedOrder.EmployeeID,
		&updatedOrder.ClientID,
		&updatedOrder.Status,
		&updatedOrder.DeliveryType,
		&updatedOrder.Total,
		&updatedOrder.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("UpdateOrder: error updating order: %w", err)
	}

	return &updatedOrder, nil
}
