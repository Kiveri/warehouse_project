package client_usecase

import (
	"time"
	"warehouse_project/internal/domain/model"
)

type (
	clientRepo interface {
		CreateClient(client *model.Client) (*model.Client, error)
	}
	timer interface {
		Now() time.Time
	}
)

type ClientUseCase struct {
	clientRepo clientRepo
	timer      timer
}

func NewClientUseCase(clientRepo clientRepo, timer timer) *ClientUseCase {
	return &ClientUseCase{
		clientRepo: clientRepo,
		timer:      timer,
	}
}
