package order_usecase

import (
	"fmt"
	"time"
	"warehouse_project/internal/domain/model"
)

func (uc *OrderUseCase) AddPositions(id []int) (*model.Order, error) {
	var positions []model.Position
	for _, id := range id {
		position, exists := uc.positionRepo.FindPosition(id)
		if exists != nil {
			return &model.Order{}, fmt.Errorf("position not found")
		}

		positions = append(positions, *position)
	}

	order := model.Order{
		Positions: positions,
		Status:    model.Created,
		DelType:   model.CourierDelivery,
		CreatedAt: time.Now(),
	}

	uc.orderRepo.CreateOrder(&order)
	return &order, nil
}
