package position_controller

import (
	"time"

	"warehouse_project/internal/domain/model"
)

type positionResponse struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	Barcode      string             `json:"barcode"`
	Price        float64            `json:"price"`
	PositionType model.PositionType `json:"position_type"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

func mapPositionToResponse(position *model.Position) *positionResponse {
	return &positionResponse{
		ID:           position.ID,
		Name:         position.Name,
		Barcode:      position.Barcode,
		Price:        position.Price,
		PositionType: position.PositionType,
		CreatedAt:    position.CreatedAt,
		UpdatedAt:    position.UpdatedAt,
	}
}
