package client_usecase

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

type FindClientReq struct {
	ID int64
}

func (cu *ClientUseCase) FindClient(req FindClientReq) (*model.Client, error) {
	client, err := cu.clientRepo.FindClient(req.ID)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.FindClient: %w", err)
	}

	return client, nil
}
