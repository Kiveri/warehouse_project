package position_usecase

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

type UpdatePositionReq struct {
	ID    int64
	Price float32
}

func (pu *PositionUseCase) UpdatePosition(req UpdatePositionReq) (*model.Position, error) {
	position, err := pu.positionRepo.FindPosition(req.ID)
	if err != nil {
		return nil, fmt.Errorf("positionRepo.FindPosition: %w", err)
	}

	position.ChangePrice(req.Price, pu.timer.Now())

	if _, err = pu.positionRepo.UpdatePosition(position); err != nil {
		return nil, fmt.Errorf("positionRepo.UpdatePosition: %w", err)
	}

	return position, nil
}
