package position_usecase

import "fmt"

type UpdatePositionReq struct {
	ID    int64
	Price float32
}

func (pu *PositionUseCase) UpdatePosition(req UpdatePositionReq) error {
	position, err := pu.positionRepo.FindPosition(req.ID)
	if err != nil {
		return fmt.Errorf("positionRepo.FindPosition: %w", err)
	}

	position.ChangePrice(req.Price)

	if err = pu.positionRepo.UpdatePosition(position); err != nil {
		return fmt.Errorf("positionRepo.UpdatePosition: %w", err)
	}

	return nil
}
