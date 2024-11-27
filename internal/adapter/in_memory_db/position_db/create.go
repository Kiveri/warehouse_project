package position_db

import "warehouse_project/internal/domain/model"

func (r *PositionRepo) CreatePosition(position *model.Position) (*model.Position, error) {
	position.ID = r.nextID
	r.positionsMap[position.ID] = position
	r.nextID++
	return position, nil
}
