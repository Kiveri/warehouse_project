package position_controller

import (
	"errors"
	"net/http"
	"strconv"

	"warehouse_project/internal/adapter/postgres/positions"
	"warehouse_project/internal/controller"
	"warehouse_project/internal/usecase/position_usecase"
)

func (c *Controller) Find(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	positionID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("position id not present", "id"))

		return
	}

	findPosition, err := c.positionUseCase.FindPosition(position_usecase.FindPositionReq{
		ID: positionID,
	})
	if err != nil {
		if errors.Is(err, positions.NotFound) {
			controller.NotFoundErrorRespond(w, controller.NewNotFoundError("position not found"))

			return
		}

		controller.InternalServerErrorRespond(w, err)

		return
	}

	controller.Respond(w, http.StatusOK, findPosition)
}
