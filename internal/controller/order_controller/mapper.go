package order_controller

import (
	"time"

	"warehouse_project/internal/domain/model"
)

type orderResponse struct {
	ID           int64                    `json:"id"`
	Positions    []*orderPositionResponse `json:"positions"`
	EmployeeID   int64                    `json:"employee_id"`
	ClientID     int64                    `json:"client_id"`
	Status       model.OrderStatus        `json:"status"`
	DeliveryType model.DeliveryType       `json:"delivery_type"`
	Total        float64                  `json:"total"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
}

func mapPositions(orderPositionMap map[int64]*model.OrderPosition) []*orderPositionResponse {
	result := make([]*orderPositionResponse, 0, len(orderPositionMap))
	for _, op := range orderPositionMap {
		result = append(result, &orderPositionResponse{
			Quantity:  op.Quantity,
			UnitPrice: op.UnitPrice,
			Position: &positionResponse{
				ID:           op.Position.ID,
				Name:         op.Position.Name,
				Barcode:      op.Position.Barcode,
				Price:        op.Position.Price,
				PositionType: op.Position.PositionType,
				CreatedAt:    op.Position.CreatedAt,
				UpdatedAt:    op.Position.UpdatedAt,
			},
		})
	}

	return result
}

func mapOrderResponse(order *model.Order) *orderResponse {
	return &orderResponse{
		ID:           order.ID,
		Positions:    mapPositions(order.Positions),
		EmployeeID:   order.EmployeeID,
		ClientID:     order.ClientID,
		Status:       order.Status,
		DeliveryType: order.DeliveryType,
		Total:        order.Total,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}

type orderPositionResponse struct {
	Quantity  int64             `json:"quantity"`
	UnitPrice float64           `json:"unit_price"`
	Position  *positionResponse `json:"position"`
}

type positionResponse struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	Barcode      string             `json:"barcode"`
	Price        float64            `json:"price"`
	PositionType model.PositionType `json:"position_type"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
