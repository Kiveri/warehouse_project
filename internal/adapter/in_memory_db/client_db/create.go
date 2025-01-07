package client_db

import "warehouse_project/internal/domain/model"

func (cr *ClientRepo) CreateClient(client *model.Client) (*model.Client, error) {
	client.ID = cr.getNextID()
	cr.clientsMap[client.ID] = client

	return client, nil
}
