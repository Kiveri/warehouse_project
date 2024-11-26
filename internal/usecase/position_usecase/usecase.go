package position_usecase

import "warehouse_project/internal/adapter/in_memory_db/position_db"

type PositionUseCase struct {
	r *position_db.PositionRepo
}

func NewPositionUseCase(r *position_db.PositionRepo) *PositionUseCase {
	return &PositionUseCase{r}
}
