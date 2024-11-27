package position_db

import (
	"fmt"
)

func (pr *PositionRepo) DeletePosition(id int64) error {
	if _, exists := pr.positionsMap[id]; !exists {
		return fmt.Errorf("position does not exist")
	}
	delete(pr.positionsMap, id)

	return nil
}
