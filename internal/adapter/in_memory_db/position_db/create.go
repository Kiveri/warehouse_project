package position_db

import "warehouse_project/internal/domain/model"

func (pr *PositionRepo) CreatePosition(position *model.Position) (*model.Position, error) {
	position.ID = pr.nextID
	pr.positionsMap[position.ID] = position
	pr.nextID++
	return position, nil
}
