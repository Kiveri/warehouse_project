package client_usecase

import "fmt"

type UpdateClientReq struct {
	ID          int64
	HomeAddress string
}

func (cu *ClientUseCase) UpdateClient(req UpdateClientReq) error {
	client, err := cu.clientRepo.FindClient(req.ID)
	if err != nil {
		return fmt.Errorf("clientRepo.FindClient: %w", err)
	}

	client.ChangeAddress(req.HomeAddress)

	if err = cu.clientRepo.UpdateClient(client); err != nil {
		return fmt.Errorf("clientRepo.UpdateClient: %w", err)
	}

	return nil
}
