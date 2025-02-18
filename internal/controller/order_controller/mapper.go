package order_controller

import (
	"time"

	"warehouse_project/internal/domain/model"
)

type orderResponse struct {
	ID           int64                          `json:"id"`
	Positions    map[int64]*model.OrderPosition `json:"positions"`
	EmployeeID   int64                          `json:"employee_id"`
	ClientID     int64                          `json:"client_id"`
	Status       model.OrderStatus              `json:"status"`
	DeliveryType model.DeliveryType             `json:"delivery_type"`
	Total        float64                        `json:"total"`
	CreatedAt    time.Time                      `json:"created_at"`
	UpdatedAt    time.Time                      `json:"updated_at"`
}

func mapOrderResponse(order *model.Order) *orderResponse {
	return &orderResponse{
		ID:           order.ID,
		Positions:    order.Positions,
		EmployeeID:   order.EmployeeID,
		ClientID:     order.ClientID,
		Status:       order.Status,
		DeliveryType: order.DeliveryType,
		Total:        order.Total,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}
