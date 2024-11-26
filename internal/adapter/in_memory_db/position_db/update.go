package position_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (r *PositionRepo) UpdatePosition(position *model.Position) error {
	if _, exists := r.positions[position.ID]; !exists {
		return fmt.Errorf("position does not exist")
	}
	r.positions[position.ID] = position
	return nil
}
