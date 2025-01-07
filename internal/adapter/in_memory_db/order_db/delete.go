package order_db

import "fmt"

func (or *OrderRepo) DeleteOrder(id int64) error {
	if _, exists := or.ordersMap[id]; !exists {
		return fmt.Errorf("order with id %d does not found", id)
	}
	delete(or.ordersMap, id)

	return nil
}
