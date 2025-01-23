package orders

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) UpdateOrder(order *model.Order) (*model.Order, error) {
	query := `
		UPDATE orders SET positions = $1, employee_id = $2, client_id = $3, status = $4, delivery_type = $5,  total = $6, updated_at = $7
        WHERE id = $8
		RETURNING id, positions, employee_id, client_id, status, delivery_type, total, updated_at
		`

	err := r.cluster.Conn.QueryRow(context.Background(), query,
		order.Positions,
		order.EmployeeID,
		order.ClientID,
		order.Status,
		order.DeliveryType,
		order.Total,
		order.UpdatedAt,
		order.ID,
	)

	if err != nil {
		return order, fmt.Errorf("UpdateOrder: error updating order: %w", err)
	}

	return order, nil
}
