package order_usecase

import "warehouse_project/internal/domain/model"

func (u *OrderUseCase) CreateOrderUC(position *model.Position) (*model.Order, error) {
	order := model.NewOrder(position)
	if err := u.r.CreateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}
