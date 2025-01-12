package orders

import (
	"context"
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) FindOne(ctx context.Context, id int64) (*model.Order, error) {
	var order model.Order
	query := "SELECT * FROM orders WHERE order_id = $1"
	err := r.cluster.Conn.QueryRow(ctx, query, id).Scan(&order.ID, &order.Positions)
	if err != nil {
		fmt.Errorf("failed to query context: %w", err)
	}

	return &order, nil
}
