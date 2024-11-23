package position_usecase

import "warehouse_project/internal/domain/model"

func (u *PositionUseCase) FindPositionUC(id int) (*model.Position, error) {
	return u.r.FindPosition(id)
}
