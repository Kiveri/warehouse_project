package position_db

import (
	"fmt"
)

func (r *PositionRepo) DeletePosition(id int) error {
	if _, exists := r.positions[id]; !exists {
		return fmt.Errorf("position does not exist")
	}
	delete(r.positions, id)
	return nil
}
