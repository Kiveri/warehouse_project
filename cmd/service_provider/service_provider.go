package service_provider

import (
	"warehouse_project/internal/adapter/in_memory_db/employee_db"
	"warehouse_project/internal/adapter/in_memory_db/order_db"
	"warehouse_project/internal/adapter/in_memory_db/position_db"
	"warehouse_project/internal/adapter/postgres/clients"
	"warehouse_project/internal/config"
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

	employeeRepo *employee_db.EmployeeRepo
	positionRepo *position_db.PositionRepo
	orderRepo    *order_db.OrderRepo
	clientRepo   *clients.Repo

	timer *timer.Timer
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
