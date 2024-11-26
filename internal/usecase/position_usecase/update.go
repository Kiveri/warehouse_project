package position_usecase

import "warehouse_project/internal/domain/model"

func (u *PositionUseCase) UpdatePosition(position *model.Position) error {
	return u.r.UpdatePosition(position)
}
