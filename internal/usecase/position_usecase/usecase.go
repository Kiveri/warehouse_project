package position_usecase

import (
	"warehouse_project/internal/domain/model"
)

type positionRepo interface {
	CreatePosition(position *model.Position) (*model.Position, error)
	DeletePosition(id int64) error
}

type PositionUseCase struct {
	positionRepo positionRepo
}

func NewPositionUseCase(positionRepo positionRepo) *PositionUseCase {
	return &PositionUseCase{
		positionRepo: positionRepo,
	}
}
