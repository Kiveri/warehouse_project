package service_provider

import (
	"net/http"
)

func (sp *ServiceProvider) GetRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /clients", sp.GetClientController().Create)
	mux.HandleFunc("GET /clients/{id}", sp.GetClientController().Find)
	mux.HandleFunc("DELETE /clients/{id}", sp.GetClientController().Delete)
	mux.HandleFunc("PUT /clients/{id}", sp.GetClientController().ChangeAddress)

	mux.HandleFunc("POST /employees", sp.GetEmployeeController().Create)
	mux.HandleFunc("GET /employees/{id}", sp.GetEmployeeController().Find)
	mux.HandleFunc("DELETE /employees/{id}", sp.GetEmployeeController().Delete)
	mux.HandleFunc("PUT /employees/{id}", sp.GetEmployeeController().ChangeRole)

	mux.HandleFunc("POST /orders", sp.GetOrderController().Create)
	mux.HandleFunc("GET /orders/{id}", sp.GetOrderController().Find)
	mux.HandleFunc("DELETE /orders/{id}", sp.GetOrderController().Delete)
	mux.HandleFunc("PUT /orders/{id}", sp.GetOrderController().ChangeStatus)

	mux.HandleFunc("POST /positions", sp.GetPositionController().Create)
	mux.HandleFunc("GET /positions/{id}", sp.GetPositionController().Find)
	mux.HandleFunc("DELETE /positions/{id}", sp.GetPositionController().Delete)
	mux.HandleFunc("PUT /positions/{id}", sp.GetPositionController().ChangePrice)

	return mux
}
