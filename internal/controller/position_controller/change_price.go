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

func (c *Controller) ChangePrice() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		positionID, err := strconv.ParseInt(id, 0, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(controller.NewValidationError("position id not present", "id"))

			return
		}

		decoder := json.NewDecoder(r.Body)
		var req changePriceRequest
		err = decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		if validationError := validateChangePriceRequest(req); validationError != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(validationError)

			return
		}

		updatePriceStatus, err := c.positionUseCase.UpdatePosition(position_usecase.UpdatePositionReq{
			ID:    positionID,
			Price: req.Price,
		})
		if err != nil {
			if errors.Is(err, positions.NotFound) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(controller.NewNotFoundError("position id not found"))

				return
			}

			http.Error(w, err.Error(), 500)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatePriceStatus)

		return
	}
}

func validateChangePriceRequest(req changePriceRequest) *controller.ValidationError {
	if req.Price <= 0 {
		return controller.NewValidationError("price must be positive", "price")
	}

	return nil
}
