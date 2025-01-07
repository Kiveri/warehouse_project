package position_db

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (pr *PositionRepo) UpdatePosition(position *model.Position) (*model.Position, error) {
	if _, exists := pr.positionsMap[position.ID]; !exists {
		return nil, fmt.Errorf("position with id %v does not exist", position.ID)
	}
	pr.positionsMap[position.ID] = position

	return position, nil
}
