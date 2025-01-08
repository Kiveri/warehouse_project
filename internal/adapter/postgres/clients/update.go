package clients

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) UpdateClient(client *model.Client) (*model.Client, error) {
	if _, exists := r.clientsMap[client.ID]; !exists {
		return nil, fmt.Errorf("employee with id %v does not exist", client.ID)
	}
	r.clientsMap[client.ID] = client

	return client, nil
}
