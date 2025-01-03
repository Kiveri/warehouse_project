package order_usecase

import "fmt"

func (ou *OrderUseCase) DeleteOrder(id int64) error {
	if err := ou.orderRepo.DeleteOrder(id); err != nil {
		return fmt.Errorf("orderRepo.DeleteOrder: %w", err)
	}

	return nil
}
