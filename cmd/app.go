package main

import (
	"net/http"

	"warehouse_project/cmd/service_provider"
)

func main() {
	sp := service_provider.NewServiceProvider()

	err := http.ListenAndServe(":8080", sp.GetRoutes())
	if err != nil {
		panic(err)
	}
}
