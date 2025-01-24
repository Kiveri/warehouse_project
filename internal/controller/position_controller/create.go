package position_controller

import (
	"encoding/json"
	"net/http"

	"warehouse_project/internal/controller"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/position_usecase"
)

type createPositionRequest struct {
	Name         string             `json:"name"`
	Barcode      string             `json:"barcode"`
	Price        float64            `json:"price"`
	PositionType model.PositionType `json:"position_type"`
}

func (c *Controller) Create() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req createPositionRequest
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		if validationError := validateCreatePositionRequest(req); validationError != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(validationError)

			return
		}

		position, err := c.positionUseCase.CreatePosition(position_usecase.CreatePositionReq{
			Name:         req.Name,
			Barcode:      req.Barcode,
			Price:        req.Price,
			PositionType: req.PositionType,
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

func validateCreatePositionRequest(req createPositionRequest) *controller.ValidationError {
	if req.Name == "" {
		return controller.NewValidationError("name is required", "name")
	}

	return nil
}
