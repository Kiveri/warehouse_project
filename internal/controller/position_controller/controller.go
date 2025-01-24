package position_controller

import (
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/position_usecase"
)

type (
	Controller struct {
		positionUseCase positionUseCase
	}

	positionUseCase interface {
		CreatePosition(req position_usecase.CreatePositionReq) (*model.Position, error)
		FindPosition(req position_usecase.FindPositionReq) (*model.Position, error)
		UpdatePosition(req position_usecase.UpdatePositionReq) (*model.Position, error)
		DeletePosition(req position_usecase.DeletePositionReq) error
	}
)

func NewController(positionUseCase positionUseCase) *Controller {
	return &Controller{
		positionUseCase: positionUseCase,
	}
}
