package clients

import (
	"fmt"

	"warehouse_project/internal/domain/model"
)

func (r *Repo) FindClient(id int64) (*model.Client, error) {
	client, exists := r.clientsMap[id]
	if !exists {
		return nil, fmt.Errorf("client not found")
	}

	return client, nil
}
