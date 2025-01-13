package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type (
	OrderStatus  int
	DeliveryType int
)

const (
	Created    OrderStatus = 1
	Building   OrderStatus = 2
	Delivering OrderStatus = 3
	Delivered  OrderStatus = 4

	CourierDelivery DeliveryType = 1
	SelfDelivery    DeliveryType = 2
	PointOfDelivery DeliveryType = 3
)

type OrderPosition struct {
	PositionID int64
	Quantity   int64
	UnitPrice  float64
	Position   *Position
}

func NewOrderPosition(positionID, quantity int64, unitPrice float64) *OrderPosition {
	return &OrderPosition{
		PositionID: positionID,
		Quantity:   quantity,
		UnitPrice:  unitPrice,
	}
}

func (op *OrderPosition) Value() (driver.Value, error) {
	if op == nil {
		return nil, nil
	}
	return json.Marshal(op)
}

func (op *OrderPosition) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return errors.New("not a string")
	}

	return json.Unmarshal([]byte(b), &op)
}

type Order struct {
	ID           int64
	Positions    []*OrderPosition
	EmployeeID   int64
	ClientID     int64
	Status       OrderStatus
	DeliveryType DeliveryType
	Total        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewOrder(employeeID, clientID int64, deliveryType DeliveryType, now time.Time) *Order {
	return &Order{
		EmployeeID:   employeeID,
		ClientID:     clientID,
		Status:       Created,
		DeliveryType: deliveryType,
		CreatedAt:    now,
	}
}

func (o *Order) ChangeStatus(newStatus OrderStatus, now time.Time) {
	o.Status = newStatus
	o.UpdatedAt = now
}
