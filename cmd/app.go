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
	orderUseCase := order_usecase.NewOrderUseCase(orderRepo, positionRepo)

	createEmployee, err := employeeUseCase.CreateEmployeeUC(employee_usecase.CreateEmployeeReq{
		Name:    "Denis",
		Surname: "Popov",
		Phone:   "79995398037",
		Email:   "denpopov.m@gmail.com",
		Post:    model.Leader,
	})
	if err != nil {
		fmt.Println(err)
	}

	createPosition, err := positionUseCase.CreatePositionUC(position_usecase.CreatePositionReq{
		Name:    "Электрический снегоуборщик Gigant SP-2300-460ES",
		Barcode: "10001",
		Price:   15349,
		PosType: model.BasicProduct,
	})

	positionRepo.CreatePosition(&model.Position{
		Name:    "Светодиодная гирлянда TDM Шишки, 50 LED, 5м, 8 режимов, многоцветная, 250 В SQ0361-0050",
		Barcode: "10002",
		Price:   557,
		PosType: model.BasicConsumable})

	order, err := orderUseCase.AddPositions([]int{1})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(createEmployee)
	fmt.Println(createPosition)
	fmt.Println(order)

}
