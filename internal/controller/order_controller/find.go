package order_controller

import (
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/orders"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/order_usecase"
)

func (c *Controller) Find(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	orderID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("order id not present", "id"))

		return
	}

	findOrder, err := c.orderUseCase.FindOrder(order_usecase.FindOrderReq{
		ID: orderID,
	})
	if err != nil {
		if errors.Is(err, orders.NotFound) {
			controller.ValidationErrorRespond(w, controller.NewValidationError("employee not found", "id"))

			return
		}

		controller.InternalServer(w, err)

		return
	}

	controller.Validation(w, http.StatusOK, findOrder)
}
