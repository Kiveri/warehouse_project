package service_provider

import (
	"warehouse_project/internal/controller/client_controller"
	"warehouse_project/internal/controller/employee_controller"
	"warehouse_project/internal/controller/order_controller"
	"warehouse_project/internal/controller/position_controller"
)

func (sp *ServiceProvider) GetClientController() *client_controller.Controller {
	if sp.clientController == nil {
		sp.clientController = client_controller.NewController(sp.GetClientUseCase())
	}

	return sp.clientController
}

func (sp *ServiceProvider) GetEmployeeController() *employee_controller.Controller {
	if sp.employeeController == nil {
		sp.employeeController = employee_controller.NewController(sp.GetEmployeeUseCase())
	}

	return sp.employeeController
}

func (sp *ServiceProvider) GetOrderController() *order_controller.Controller {
	if sp.orderController == nil {
		sp.orderController = order_controller.NewController(sp.GetOrderUseCase())
	}

	return sp.orderController
}

func (sp *ServiceProvider) GetPositionController() *position_controller.Controller {
	if sp.positionController == nil {
		sp.positionController = position_controller.NewController(sp.GetPositionUseCase())
	}

	return sp.positionController
}
