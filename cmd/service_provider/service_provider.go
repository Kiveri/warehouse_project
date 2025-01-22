package service_provider

import (
	"warehouse_project/internal/adapter/postgres/clients"
	"warehouse_project/internal/adapter/postgres/employees"
	"warehouse_project/internal/adapter/postgres/orders"
	"warehouse_project/internal/adapter/postgres/positions"
	"warehouse_project/internal/config"
	"warehouse_project/internal/controller/client_controller"
	"warehouse_project/internal/pkg/timer"
	"warehouse_project/internal/usecase/client_usecase"
	"warehouse_project/internal/usecase/employee_usecase"
	"warehouse_project/internal/usecase/order_usecase"
	"warehouse_project/internal/usecase/position_usecase"
)

type ServiceProvider struct {
	dbCluster *config.Cluster

	employeeUseCase *employee_usecase.EmployeeUseCase
	positionUseCase *position_usecase.PositionUseCase
	orderUseCase    *order_usecase.OrderUseCase
	clientUseCase   *client_usecase.ClientUseCase

	employeeRepo *employees.Repo
	positionRepo *positions.Repo
	orderRepo    *orders.Repo
	clientRepo   *clients.Repo

	clientController *client_controller.Controller

	timer *timer.Timer
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
