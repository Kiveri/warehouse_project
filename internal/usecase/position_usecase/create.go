package position_usecase

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

type CreatePositionReq struct {
	Name    string
	Barcode string
	Price   float64
	PosType model.PositionType
}

func (pu *PositionUseCase) CreatePosition(req CreatePositionReq) (*model.Position, error) {
	now := pu.timer.Now()
	position := model.NewPosition(req.Name, req.Barcode, req.Price, req.PosType, now)
	position, err := pu.positionRepo.CreatePosition(position)
	if err != nil {
		return nil, fmt.Errorf("positionRepo.CreatePosition: %w", err)
	}

	return position, nil
}
