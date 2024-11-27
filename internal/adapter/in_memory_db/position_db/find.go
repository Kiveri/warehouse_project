package position_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (r *PositionRepo) FindPosition(id int) (*model.Position, error) {
	position, exists := r.positionsMap[id]
	if !exists {
		return nil, fmt.Errorf("position not found")
	}
	return position, nil
}
