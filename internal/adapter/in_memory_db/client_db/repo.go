package client_db

import "warehouse_project/internal/domain/model"

type ClientRepo struct {
	clientsMap map[int64]*model.Client
	nextID     int64
}

func (cr *ClientRepo) getNextID() int64 {
	nextID := cr.nextID
	cr.nextID++

	return nextID
}

func NewClientRepo() *ClientRepo {
	return &ClientRepo{
		clientsMap: make(map[int64]*model.Client),
		nextID:     1,
	}
}
