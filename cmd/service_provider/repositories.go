package service_provider

import (
	"context"

	"warehouse_project/internal/adapter/in_memory_db/employee_db"
	"warehouse_project/internal/adapter/in_memory_db/order_db"
	"warehouse_project/internal/adapter/in_memory_db/position_db"
	"warehouse_project/internal/adapter/postgres/clients"
)

func (sp *ServiceProvider) getEmployeeRepository() *employee_db.EmployeeRepo {
	if sp.employeeRepo == nil {
		sp.employeeRepo = employee_db.NewEmployeeRepo()
	}

	return sp.employeeRepo
}

func (sp *ServiceProvider) getPositionRepository() *position_db.PositionRepo {
	if sp.positionRepo == nil {
		sp.positionRepo = position_db.NewPositionRepo()
	}

	return sp.positionRepo
}

func (sp *ServiceProvider) getOrderRepository() *order_db.OrderRepo {
	if sp.orderRepo == nil {
		sp.orderRepo = order_db.NewOrderRepo()
	}

	return sp.orderRepo
}

func (sp *ServiceProvider) getClientRepository() *clients.Repo {
	if sp.clientRepo == nil {
		sp.clientRepo = clients.NewRepo(sp.getDbCluster(context.Background()))
	}

	return sp.clientRepo
}
