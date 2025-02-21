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

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req createPositionRequest
	err := decoder.Decode(&req)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("position id not present", "id"))

		return
	}

	if validationError := validateCreatePositionRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	position, err := c.positionUseCase.CreatePosition(position_usecase.CreatePositionReq{
		Name:         req.Name,
		Barcode:      req.Barcode,
		Price:        req.Price,
		PositionType: req.PositionType,
	})
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}

	if err = controller.EncodeResponse(w, mapPositionToResponse(position)); err != nil {
		return
	}
}

func validateCreatePositionRequest(req createPositionRequest) *controller.ValidationError {
	if req.Name == "" {
		return controller.NewValidationError("name is required", "name")
	}
	if req.Price <= 0 {
		return controller.NewValidationError("price must be positive", "price")
	}
	if req.PositionType <= 0 || req.PositionType >= 5 {
		return controller.NewValidationError("position type 1 - BasicProduct, 2 - BasicConsumable 3 - Liquid, 4 - OversizeProduct",
			"position_type")
	}

	return nil
}
