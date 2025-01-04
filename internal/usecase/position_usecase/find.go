package position_usecase

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

type FindPositionReq struct {
	ID int64
}

func (pu *PositionUseCase) FindPosition(req FindPositionReq) (*model.Position, error) {
	position, err := pu.positionRepo.FindPosition(req.ID)
	if err != nil {
		return nil, fmt.Errorf("positionRepo.FindPosition: %w", err)
	}

	return position, nil
}
