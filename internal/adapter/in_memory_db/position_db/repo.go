package position_db

import "warehouse_project/internal/domain/model"

type PositionRepo struct {
	positionsMap map[int]*model.Position
	nextID       int
}

func NewPositionRepo() *PositionRepo {
	return &PositionRepo{
		positionsMap: make(map[int]*model.Position),
		nextID:       1,
	}
}
