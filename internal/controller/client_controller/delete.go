package client_controller

import (
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/clients"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/client_usecase"
)

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	clientID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("client id not present", "id"))

		return
	}

	err = c.clientUseCase.DeleteClient(client_usecase.DeleteClientReq{
		ID: clientID,
	})
	if err != nil {
		if errors.Is(err, clients.NotFound) {
			controller.ValidationErrorRespond(w, controller.NewValidationError("client not found", "id"))

			return
		}

		controller.InternalServer(w, err)
	}
}
