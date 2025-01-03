package client_usecase

import "fmt"

func (cu *ClientUseCase) DeleteClient(id int64) error {
	if err := cu.clientRepo.DeleteClient(id); err != nil {
		return fmt.Errorf("clientRepo.DeleteClient: %w", err)
	}

	return nil
}
