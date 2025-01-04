package client_usecase

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

type UpdateClientReq struct {
	ID          int64
	HomeAddress string
}

func (cu *ClientUseCase) UpdateClient(req UpdateClientReq) (*model.Client, error) {
	client, err := cu.clientRepo.FindClient(req.ID)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.FindClient: %w", err)
	}

	client.ChangeAddress(req.HomeAddress, cu.timer.Now())

	if _, err = cu.clientRepo.UpdateClient(client); err != nil {
		return nil, fmt.Errorf("clientRepo.UpdateClient: %w", err)
	}

	return client, nil
}
