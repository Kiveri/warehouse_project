package service_provider

import (
	"context"

	"warehouse_project/internal/adapter/postgres/clients"
	"warehouse_project/internal/adapter/postgres/employees"
	"warehouse_project/internal/adapter/postgres/orders"
	"warehouse_project/internal/adapter/postgres/positions"
)

func (sp *ServiceProvider) getEmployeeRepository() *employees.Repo {
	if sp.employeeRepo == nil {
		sp.employeeRepo = employees.NewRepo(sp.getDbCluster(context.Background()))
	}

	return sp.employeeRepo
}

func (sp *ServiceProvider) getPositionRepository() *positions.Repo {
	if sp.positionRepo == nil {
		sp.positionRepo = positions.NewRepo(sp.getDbCluster(context.Background()))
	}

	return sp.positionRepo
}

func (sp *ServiceProvider) getOrderRepository() *orders.Repo {
	if sp.orderRepo == nil {
		sp.orderRepo = orders.NewRepo(sp.getDbCluster(context.Background()))
	}

	return sp.orderRepo
}

func (sp *ServiceProvider) getClientRepository() *clients.Repo {
	if sp.clientRepo == nil {
		sp.clientRepo = clients.NewRepo(sp.getDbCluster(context.Background()))
	}

	return sp.clientRepo
}
