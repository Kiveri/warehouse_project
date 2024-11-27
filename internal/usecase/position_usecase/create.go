package position_usecase

import (
	"time"
	"warehouse_project/internal/domain/model"
)

type CreatePositionReq struct {
	Name    string
	Barcode string
	Price   float32
	PosType model.PositionType
}

func (pu *PositionUseCase) CreatePosition(req CreatePositionReq) (*model.Position, error) {
	now := time.Now()
	position := model.NewPosition(req.Name, req.Barcode, req.Price, req.PosType, now)
	position, err := pu.pr.CreatePosition(position)
	if err != nil {
		return nil, err
	}
	return position, nil
}
