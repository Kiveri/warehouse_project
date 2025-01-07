package client_db

import (
	"fmt"
	"warehouse_project/internal/domain/model"
)

func (cr *ClientRepo) FindClient(id int64) (*model.Client, error) {
	client, exists := cr.clientsMap[id]
	if !exists {
		return nil, fmt.Errorf("client not found")
	}

	return client, nil
}
