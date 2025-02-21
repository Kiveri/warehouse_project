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

func (c *Controller) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	orderID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("order id not present", "id"))

		return
	}

	decoder := json.NewDecoder(r.Body)
	var req changeStatusRequest
	err = decoder.Decode(&req)
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}

	if validationError := validateChangeStatusRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	updateOrderStatus, err := c.orderUseCase.UpdateOrder(order_usecase.UpdateOrderReq{
		ID:     orderID,
		Status: req.Status,
	})
	if err != nil {
		if errors.Is(err, orders.NotFound) {
			controller.NotFoundErrorRespond(w, controller.NewNotFoundError("order not found"))

			return
		}

		controller.InternalServerErrorRespond(w, err)

		return
	}

	if err = controller.EncodeResponse(w, mapOrderResponse(updateOrderStatus)); err != nil {
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
