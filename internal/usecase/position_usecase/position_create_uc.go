package position_usecase

import (
	"fmt"
	"time"
	"warehouse_project/internal/domain/model"
)

type CreatePositionReq struct {
	Name    string
	Barcode string
	Price   float32
	PosType model.PositionType
}

func (u *PositionUseCase) CreatePositionUC(req CreatePositionReq) (*model.Position, error) {
	now := time.Now()
	position := model.NewPosition(req.Name, req.Barcode, req.Price, req.PosType, now)
	position, err := u.r.CreatePosition(position)
	if err != nil {
		return nil, fmt.Errorf("error creating position: %w", err)
	}

	return position, nil
}
