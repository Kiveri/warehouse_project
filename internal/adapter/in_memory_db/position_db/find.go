package position_db

import "warehouse_project/internal/domain/model"

func (pr *PositionRepo) FindPosition(id int64) (*model.Position, bool) {
	product, found := pr.positionsMap[id]
	return product, found
}
