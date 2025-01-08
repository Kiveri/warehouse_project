package service_provider

import (
	"warehouse_project/internal/usecase/client_usecase"
	"warehouse_project/internal/usecase/employee_usecase"
	"warehouse_project/internal/usecase/order_usecase"
	"warehouse_project/internal/usecase/position_usecase"
)

func (sp *ServiceProvider) GetEmployeeUseCase() *employee_usecase.EmployeeUseCase {
	if sp.employeeUseCase == nil {
		sp.employeeUseCase = employee_usecase.NewEmployeeUseCase(sp.getEmployeeRepository(), sp.getTimer())
	}

	return sp.employeeUseCase
}

func (sp *ServiceProvider) GetPositionUseCase() *position_usecase.PositionUseCase {
	if sp.positionUseCase == nil {
		sp.positionUseCase = position_usecase.NewPositionUseCase(sp.getPositionRepository(), sp.getTimer())
	}

	return sp.positionUseCase
}

func (sp *ServiceProvider) GetOrderUseCase() *order_usecase.OrderUseCase {
	if sp.orderUseCase == nil {
		sp.orderUseCase = order_usecase.NewOrderUseCase(sp.getOrderRepository(), sp.getPositionRepository(), sp.getEmployeeRepository(), sp.getClientRepository(), sp.getTimer())
	}

	return sp.orderUseCase
}

func (sp *ServiceProvider) GetClientUseCase() *client_usecase.ClientUseCase {
	if sp.clientUseCase == nil {
		sp.clientUseCase = client_usecase.NewClientUseCase(sp.getClientRepository(), sp.getTimer())
	}

	return sp.clientUseCase
}
