package position_db

import (
	"fmt"
)

func (pr *PositionRepo) DeletePosition(id int64) error {
	if _, exists := pr.positionsMap[id]; !exists {
		return fmt.Errorf("position with id %d does not found", id)
	}
	delete(pr.positionsMap, id)

	return nil
}
