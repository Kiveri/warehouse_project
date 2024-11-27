package position_db

import "warehouse_project/internal/domain/model"

func (pr *PositionRepo) CreatePosition(position *model.Position) (*model.Position, error) {
	position.ID = pr.getNextID()
	pr.positionsMap[position.ID] = position

	return position, nil
}
