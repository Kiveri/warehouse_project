package main

import (
	"fmt"
	"warehouse_project/internal/adapter/in_memory_db/employee_db"
	"warehouse_project/internal/adapter/in_memory_db/order_db"
	"warehouse_project/internal/adapter/in_memory_db/position_db"
	"warehouse_project/internal/domain/model"
	"warehouse_project/internal/usecase/employee_usecase"
	"warehouse_project/internal/usecase/order_usecase"
	"warehouse_project/internal/usecase/position_usecase"
)

func main() {
	employeeRepo := employee_db.NewEmployeeRepo()
	positionRepo := position_db.NewPositionRepo()
	orderRepo := order_db.NewOrderRepo()

	employeeUseCase := employee_usecase.NewEmployeeUseCase(employeeRepo)
	positionUseCase := position_usecase.NewPositionUseCase(positionRepo)
	orderUseCase := order_usecase.NewOrderUseCase(orderRepo, positionRepo, employeeRepo)

	createEmployee1, err := employeeUseCase.CreateEmployee(employee_usecase.CreateEmployeeReq{
		Name:    "Denis",
		Surname: "Popov",
		Phone:   "79995398037",
		Email:   "denpopov.m@gmail.com",
		Role:    model.Manager,
	})
	if err != nil {
		fmt.Println(err)
	}

	createPosition1, err := positionUseCase.CreatePosition(position_usecase.CreatePositionReq{
		Name:    "Электрический снегоуборщик Gigant SP-2300-460ES",
		Barcode: "10001",
		Price:   15349,
		PosType: model.BasicProduct,
	})
	if err != nil {
		fmt.Println(err)
	}

	createPosition2, err := positionUseCase.CreatePosition(position_usecase.CreatePositionReq{
		Name:    "Светодиодная гирлянда TDM Шишки, 50 LED, 5м, 8 режимов, многоцветная, 250 В SQ0361-0050",
		Barcode: "10002",
		Price:   557,
		PosType: model.BasicConsumable,
	})
	if err != nil {
		fmt.Println(err)
	}

	createPosition3, err := positionUseCase.CreatePosition(position_usecase.CreatePositionReq{
		Name:    "Промывка двигателя LAVR 5-минутная классическая, 345 мл Ln1003N",
		Barcode: "10003",
		Price:   297,
		PosType: model.Liquid,
	})
	if err != nil {
		fmt.Println(err)
	}

	orderReq := order_usecase.CreateOrderReq{
		EmployeeID:   createEmployee1.ID,
		PositionsIDs: []int64{createPosition1.ID, createPosition2.ID},
		DeliveryType: model.CourierDelivery,
	}

	createOrder1, err := orderUseCase.CreateOrder(orderReq)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(createEmployee1)
	fmt.Println(createPosition1, createPosition2, createPosition3)
	fmt.Println(createOrder1)
}
