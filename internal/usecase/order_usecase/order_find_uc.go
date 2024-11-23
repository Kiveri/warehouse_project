package order_usecase

import "warehouse_project/internal/domain/model"

func (u *OrderUseCase) FindOrderUC(id int) (*model.Order, error) {
	return u.r.FindOrder(id)
}
