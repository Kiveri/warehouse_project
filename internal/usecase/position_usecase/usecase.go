package position_usecase

import (
	"time"
	"warehouse_project/internal/domain/model"
)

type PositionUseCase struct {
	positionRepo positionRepo
	timer        timer
}

type positionRepo interface {
	CreatePosition(position *model.Position) (*model.Position, error)
	DeletePosition(id int64) error
	UpdatePosition(position *model.Position) error
	FindPosition(id int64) (*model.Position, error)
}

type timer interface {
	Now() time.Time
}

func NewPositionUseCase(positionRepo positionRepo, timer timer) *PositionUseCase {
	return &PositionUseCase{
		positionRepo: positionRepo,
		timer:        timer,
	}
}
