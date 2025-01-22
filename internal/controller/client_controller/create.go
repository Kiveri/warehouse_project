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

func (c *Controller) Create() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req createClientRequest
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		if validationError := validateCreateClientRequest(req); validationError != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(validationError)

			return
		}

		client, err := c.clientUseCase.CreateClient(client_usecase.CreateClientReq{
			Name:        req.Name,
			Phone:       req.Phone,
			Email:       req.Email,
			HomeAddress: req.HomeAddress,
		})
		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(client)

		return
	}
}

func validateCreateClientRequest(req createClientRequest) *controller.ValidationError {
	if req.Name == "" {
		return controller.NewValidationError("name is required", "name")
	}

	return nil
}
