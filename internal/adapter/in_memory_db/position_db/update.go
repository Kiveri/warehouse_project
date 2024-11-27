package position_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (r *PositionRepo) UpdatePosition(position *model.Position) error {
	if _, exists := r.positionsMap[position.ID]; !exists {
		return fmt.Errorf("position does not exist")
	}
	r.positionsMap[position.ID] = position
	return nil
}
