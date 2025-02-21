package position_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/positions"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/position_usecase"
)

type changePriceRequest struct {
	Price float64 `json:"price"`
}

func (c *Controller) ChangePrice(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	positionID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("position id not present", "id"))

		return
	}

	decoder := json.NewDecoder(r.Body)
	var req changePriceRequest
	err = decoder.Decode(&req)
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}

	if validationError := validateChangePriceRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	updatePrice, err := c.positionUseCase.UpdatePosition(position_usecase.UpdatePositionReq{
		ID:    positionID,
		Price: req.Price,
	})
	if err != nil {
		if errors.Is(err, positions.NotFound) {
			controller.NotFoundErrorRespond(w, controller.NewNotFoundError("position not found"))

			return
		}

		controller.InternalServerErrorRespond(w, err)

		return
	}

	if err = controller.EncodeResponse(w, mapPositionToResponse(updatePrice)); err != nil {
		return
	}
}

func validateChangePriceRequest(req changePriceRequest) *controller.ValidationError {
	if req.Price <= 0 {
		return controller.NewValidationError("price must be positive", "price")
	}

	return nil
}
