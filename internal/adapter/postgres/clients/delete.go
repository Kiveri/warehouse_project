package clients

import "fmt"

func (r *Repo) DeleteClient(id int64) error {
	if _, exists := r.clientsMap[id]; !exists {
		return fmt.Errorf("client with id %d does not found", id)
	}
	delete(r.clientsMap, id)

	return nil
}
