package position_usecase

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (pu *PositionUseCase) FindPosition(id int64) (*model.Position, error) {
	position, err := pu.positionRepo.FindPosition(id)
	if err != nil {
		return nil, fmt.Errorf("positionRepo.FindPosition: %w", err)
	}

	return position, nil
}
