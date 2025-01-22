package orders

import (
	"context"
	"errors"
	"fmt"

	"warehouse_project/internal/domain/model"

	"github.com/jackc/pgx/v5"
)

func (r *Repo) FindOrder(id int64) (*model.Order, error) {
	var order model.Order

	query := `
		SELECT * FROM orders
		WHERE id = $1;
	`

	err := r.cluster.Conn.QueryRow(context.Background(), query, id).Scan(
		&order.ID,
		&order.Positions,
		&order.EmployeeID,
		&order.ClientID,
		&order.Status,
		&order.DeliveryType,
		&order.Total,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("order with id %d not found", id)
		}
		return nil, fmt.Errorf("error finding order with id %d: %w", id, err)
	}

	return &order, nil
}
