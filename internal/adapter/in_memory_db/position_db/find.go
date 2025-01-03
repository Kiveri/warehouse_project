package position_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (pr *PositionRepo) FindPosition(id int64) (*model.Position, error) {
	position, exists := pr.positionsMap[id]
	if !exists {
		return nil, fmt.Errorf("position with id %d does not found", id)
	}

	return position, nil
}
