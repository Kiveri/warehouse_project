package service_provider

import (
	"warehouse_project/internal/controller/client_controller"
)

func (sp *ServiceProvider) GetClientController() *client_controller.Controller {
	if sp.clientController == nil {
		sp.clientController = client_controller.NewController(sp.GetClientUseCase())
	}

	return sp.clientController
}
