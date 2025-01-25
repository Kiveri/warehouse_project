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

func (c *Controller) ChangeAddress() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		clientID, err := strconv.ParseInt(id, 0, 64)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(controller.NewValidationError("client id not present", "id"))

			return
		}

		decoder := json.NewDecoder(r.Body)
		var req changeAddressRequest
		err = decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		if validationError := validateChangeAddressRequest(req); validationError != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(validationError)

			return
		}

		updateClientAddress, err := c.clientUseCase.UpdateClient(client_usecase.UpdateClientReq{
			ID:          clientID,
			HomeAddress: req.HomeAddress,
		})
		if err != nil {
			if errors.Is(err, clients.NotFound) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(controller.NewNotFoundError("client id not found"))

				return
			}

			http.Error(w, err.Error(), 500)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updateClientAddress)

		return
	}
}

func validateChangeAddressRequest(req changeAddressRequest) *controller.ValidationError {
	if req.HomeAddress == "" {
		return controller.NewValidationError("home address is required", "home_address")
	}

	return nil
}
