package position_db

import (
	"fmt"
)

func (r *PositionRepo) DeletePosition(id int) error {
	if _, exists := r.positionsMap[id]; !exists {
		return fmt.Errorf("position does not exist")
	}
	delete(r.positionsMap, id)
	return nil
}
