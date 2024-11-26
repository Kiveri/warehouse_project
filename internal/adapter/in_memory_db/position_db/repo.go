package position_db

import "warehouse_project/internal/domain/model"

type PositionRepo struct {
	positions map[int]*model.Position
	nextID    int
}

func NewPositionRepo() *PositionRepo {
	return &PositionRepo{
		positions: make(map[int]*model.Position),
		nextID:    1,
	}
}
