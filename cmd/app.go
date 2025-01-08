package main

import (
	"fmt"

	"warehouse_project/cmd/service_provider"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/client_usecase"
	"warehouse_project/internal/usecase/employee_usecase"
	"warehouse_project/internal/usecase/order_usecase"
	"warehouse_project/internal/usecase/position_usecase"
)

func main() {
	sp := service_provider.NewServiceProvider()

	createEmployee1, err := sp.GetEmployeeUseCase().CreateEmployee(employee_usecase.CreateEmployeeReq{
		Name:  "Denis Popov",
		Phone: "79995398037",
		Email: "denpopov.m@gmail.com",
		Role:  model.Manager,
	})
	if err != nil {
		fmt.Println(err)
	}

	createPosition1, err := sp.GetPositionUseCase().CreatePosition(position_usecase.CreatePositionReq{
		Name:    "Электрический снегоуборщик Gigant SP-2300-460ES",
		Barcode: "10001",
		Price:   15349,
		PosType: model.BasicProduct,
	})
	if err != nil {
		fmt.Println(err)
	}

	createPosition2, err := sp.GetPositionUseCase().CreatePosition(position_usecase.CreatePositionReq{
		Name:    "Светодиодная гирлянда TDM Шишки, 50 LED, 5м, 8 режимов, многоцветная, 250 В SQ0361-0050",
		Barcode: "10002",
		Price:   557,
		PosType: model.BasicConsumable,
	})
	if err != nil {
		fmt.Println(err)
	}

	createPosition3, err := sp.GetPositionUseCase().CreatePosition(position_usecase.CreatePositionReq{
		Name:    "Промывка двигателя LAVR 5-минутная классическая, 345 мл Ln1003N",
		Barcode: "10003",
		Price:   297,
		PosType: model.Liquid,
	})
	if err != nil {
		fmt.Println(err)
	}

	clientReq := client_usecase.CreateClientReq{
		Name:        "Nikita Popkov",
		Phone:       "89991234567",
		Email:       "nikitapopkov@gmail.com",
		HomeAddress: "SPB",
	}

	createClient1, err := sp.GetClientUseCase().CreateClient(clientReq)
	if err != nil {
		fmt.Println(err)
	}

	orderReq := order_usecase.CreateOrderReq{
		EmployeeID:   createEmployee1.ID,
		PositionsIDs: []int64{createPosition1.ID, createPosition2.ID},
		DeliveryType: model.CourierDelivery,
		ClientID:     createClient1.ID,
	}

	createOrder1, err := sp.GetOrderUseCase().CreateOrder(orderReq)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(createEmployee1)
	fmt.Println(createPosition1, createPosition2, createPosition3)
	fmt.Println(createOrder1)
	fmt.Println(createClient1)
}
