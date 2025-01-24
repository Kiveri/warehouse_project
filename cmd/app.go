package main

import (
	"net/http"

	"warehouse_project/cmd/service_provider"
)

func main() {
	sp := service_provider.NewServiceProvider()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /clients", sp.GetClientController().Create())
	mux.HandleFunc("GET /clients/{id}", sp.GetClientController().Find())

	mux.HandleFunc("POST /employees", sp.GetEmployeeController().Create())
	mux.HandleFunc("GET /employees/{id}", sp.GetEmployeeController().Find())

	mux.HandleFunc("POST /orders", sp.GetOrderController().Create())
	mux.HandleFunc("GET /orders/{id}", sp.GetOrderController().Find())

	mux.HandleFunc("POST /positions", sp.GetPositionController().Create())
	mux.HandleFunc("GET /positions/{id}", sp.GetPositionController().Find())

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
