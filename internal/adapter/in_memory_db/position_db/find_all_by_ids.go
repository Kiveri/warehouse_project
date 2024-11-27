package position_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (pr *PositionRepo) FindAllByIDs(id []int64) ([]*model.Position, error) {
	positions := make([]*model.Position, 0, len(id))

	for _, positionsId := range id {
		if position, exists := pr.positionsMap[positionsId]; exists {
			positions = append(positions, position)
		} else {
			return nil, fmt.Errorf("position with id %d not found", positionsId)
		}
	}

	return positions, nil
}
