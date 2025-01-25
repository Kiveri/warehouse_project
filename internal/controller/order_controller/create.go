package order_controller

import (
	"encoding/json"
	"net/http"

	"warehouse_project/internal/controller"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/order_usecase"
)

type createOrderRequest struct {
	PositionIDs  []int64            `json:"position_ids"`
	EmployeeID   int64              `json:"employee_id"`
	ClientID     int64              `json:"client_id"`
	Status       model.OrderStatus  `json:"status"`
	DeliveryType model.DeliveryType `json:"delivery_type"`
}

func (c *Controller) Create() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req createOrderRequest
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		if validationError := validateCreateOrderRequest(req); validationError != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(validationError)

			return
		}

		position, err := c.orderUseCase.CreateOrder(order_usecase.CreateOrderReq{
			PositionIDs:  req.PositionIDs,
			EmployeeID:   req.EmployeeID,
			ClientID:     req.ClientID,
			Status:       req.Status,
			DeliveryType: req.DeliveryType,
		})
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(position)

		return
	}
}

func validateCreateOrderRequest(req createOrderRequest) *controller.ValidationError {
	if req.PositionIDs == nil {
		return controller.NewValidationError("positions is required", "position_ids")
	}
	if req.Status != 1 {
		return controller.NewValidationError("status must be 1 - created", "status")
	}

	return nil
}
