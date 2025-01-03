package client_usecase

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

type CreateClientReq struct {
	Name        string
	Phone       string
	Email       string
	HomeAddress string
}

func (cu *ClientUseCase) CreateClient(req CreateClientReq) (*model.Client, error) {
	now := cu.timer.Now()

	client := model.NewClient(req.Name, req.Phone, req.Email, req.HomeAddress, now)
	client, err := cu.clientRepo.CreateClient(client)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.CreateClient: %w", err)
	}

	return client, nil
}
