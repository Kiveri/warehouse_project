package position_usecase

import "fmt"

type DeletePositionReq struct {
	ID int64
}

func (pu *PositionUseCase) DeletePosition(req DeletePositionReq) error {
	if err := pu.positionRepo.DeletePosition(req.ID); err != nil {
		return fmt.Errorf("positionRepo.DeletePosition: %w", err)
	}

	return nil
}
