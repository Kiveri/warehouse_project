package client_db

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (cr *ClientRepo) UpdateClient(client *model.Client) (*model.Client, error) {
	if _, exists := cr.clientsMap[client.ID]; !exists {
		return nil, fmt.Errorf("employee with id %v does not exist", client.ID)
	}
	cr.clientsMap[client.ID] = client

	return client, nil
}
