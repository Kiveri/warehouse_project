package client_db

import "fmt"

func (cr *ClientRepo) DeleteClient(id int64) error {
	if _, exists := cr.clientsMap[id]; !exists {
		return fmt.Errorf("client with id %d does not found", id)
	}
	delete(cr.clientsMap, id)

	return nil
}
