package client_controller

import (
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/clients"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/client_usecase"
)

func (c *Controller) Find(w http.ResponseWriter, r *http.Request) {
	// Извлечение ID из пути
	id := r.PathValue("id")
	clientID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("client id not present", "id"))

		return
	}

	// Поиск клиента
	findClient, err := c.clientUseCase.FindClient(client_usecase.FindClientReq{
		ID: clientID,
	})
	if err != nil {
		if errors.Is(err, clients.NotFound) {
			controller.NotFoundErrorRespond(w, controller.NewNotFoundError("client not found"))

			return
		}

		controller.InternalServerErrorRespond(w, err)

		return
	}

	// Успешный ответ
	if err = controller.EncodeResponse(w, mapClientToResponse(findClient)); err != nil {
		return
	}
}
