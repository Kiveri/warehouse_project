package position_usecase

import "warehouse_project/internal/adapter/in_memory_db/position_db"

type PositionUseCase struct {
	positionRepo *position_db.PositionRepo
}

func NewPositionUseCase(positionRepo *position_db.PositionRepo) *PositionUseCase {
	return &PositionUseCase{
		positionRepo: positionRepo,
	}
}
