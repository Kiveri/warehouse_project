package position_db

import "warehouse_project/internal/domain/model"

type PositionRepo struct {
	positionsMap map[int64]*model.Position
	nextID       int64
}

func (pr *PositionRepo) getNextID() int64 {
	nextID := pr.nextID
	pr.nextID++

	return nextID
}

func NewPositionRepo() *PositionRepo {
	return &PositionRepo{
		positionsMap: make(map[int64]*model.Position),
		nextID:       1,
	}
}
