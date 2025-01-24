package client_controller

import (
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/client_usecase"
)

type (
	Controller struct {
		clientUseCase clientUseCase
	}

	clientUseCase interface {
		CreateClient(req client_usecase.CreateClientReq) (*model.Client, error)
		UpdateClient(req client_usecase.UpdateClientReq) (*model.Client, error)
		FindClient(req client_usecase.FindClientReq) (*model.Client, error)
	}
)

func NewController(clientUseCase clientUseCase) *Controller {
	return &Controller{
		clientUseCase: clientUseCase,
	}
}
