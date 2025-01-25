package order_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/orders"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/order_usecase"
)

type changeStatusRequest struct {
	Status model.OrderStatus `json:"status"`
}

func (c *Controller) ChangeStatus() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		orderID, err := strconv.ParseInt(id, 0, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(controller.NewValidationError("order id not present", "id"))

			return
		}

		decoder := json.NewDecoder(r.Body)
		var req changeStatusRequest
		err = decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		if validationError := validateChangeStatusRequest(req); validationError != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(validationError)

			return
		}

		updateOrderStatus, err := c.orderUseCase.UpdateOrder(order_usecase.UpdateOrderReq{
			ID:     orderID,
			Status: req.Status,
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
		json.NewEncoder(w).Encode(updateOrderStatus)

		return
	}
}

func validateChangeStatusRequest(req changeStatusRequest) *controller.ValidationError {
	if req.Status <= 0 || req.Status >= 5 {
		return controller.NewValidationError("status 1 - created, 2 - building, 3 - delivering, 4 - delivered",
			"status")
	}

	return nil
}
