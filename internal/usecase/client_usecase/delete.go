package client_usecase

import "fmt"

type DeleteClientReq struct {
	ID int64
}

func (cu *ClientUseCase) DeleteClient(req DeleteClientReq) error {
	if _, err := cu.clientRepo.DeleteClient(req.ID); err != nil {
		return fmt.Errorf("clientRepo.DeleteClient: %w", err)
	}

	return nil
}
