package client_usecase

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (cu *ClientUseCase) FindClient(id int64) (*model.Client, error) {
	client, err := cu.clientRepo.FindClient(id)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.FindClient: %w", err)
	}

	return client, nil
}
