package client_controller

import (
	"encoding/json"
	"net/http"

	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/client_usecase"
)

type createClientRequest struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	HomeAddress string `json:"home_address"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req createClientRequest
	err := decoder.Decode(&req)
	if err != nil {
		controller.InternalServer(w, err)

		return
	}

	if validationError := validateCreateClientRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	client, err := c.clientUseCase.CreateClient(client_usecase.CreateClientReq{
		Name:        req.Name,
		Phone:       req.Phone,
		Email:       req.Email,
		HomeAddress: req.HomeAddress,
	})
	if err != nil {
		controller.InternalServer(w, err)

		return
	}

	controller.Validation(w, http.StatusOK, client)
}

func validateCreateClientRequest(req createClientRequest) *controller.ValidationError {
	if req.Name == "" {
		return controller.NewValidationError("name is required", "name")
	}

	return nil
}
