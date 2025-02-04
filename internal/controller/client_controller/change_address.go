package client_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/clients"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/client_usecase"
)

type changeAddressRequest struct {
	HomeAddress string `json:"home_address"`
}

func (c *Controller) ChangeAddress(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	clientID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("client id not present", "id"))

		return
	}

	decoder := json.NewDecoder(r.Body)
	var req changeAddressRequest
	err = decoder.Decode(&req)
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}

	if validationError := validateChangeAddressRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	updateClientAddress, err := c.clientUseCase.UpdateClient(client_usecase.UpdateClientReq{
		ID:          clientID,
		HomeAddress: req.HomeAddress,
	})
	if err != nil {
		if errors.Is(err, clients.NotFound) {
			controller.NotFoundErrorRespond(w, controller.NewNotFoundError("client not found"))

			return
		}

		controller.InternalServerErrorRespond(w, err)

		return
	}

	controller.Respond(w, http.StatusOK, updateClientAddress)
}

func validateChangeAddressRequest(req changeAddressRequest) *controller.ValidationError {
	if req.HomeAddress == "" {
		return controller.NewValidationError("home address is required", "home_address")
	}

	return nil
}
