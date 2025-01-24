package order_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/orders"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/order_usecase"
)

func (c *Controller) Find() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		orderID, err := strconv.ParseInt(id, 0, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(controller.NewValidationError("order id not present", "id"))

			return
		}

		findOrder, err := c.orderUseCase.FindOrder(order_usecase.FindOrderReq{
			ID: orderID,
		})
		if err != nil {
			if errors.Is(err, orders.NotFound) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(controller.NewNotFoundError("order id not found"))

				return
			}

			http.Error(w, err.Error(), 500)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(findOrder)

		return
	}
}
