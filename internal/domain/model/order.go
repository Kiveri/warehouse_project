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
	ID         int64
	PositionID int64
	Quantity   int64
	UnitPrice  float64
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

func NewOrder(positions []*OrderPosition, employeeID, clientID int64, deliveryType DeliveryType, total float64, now time.Time) *Order {
	return &Order{
		Positions:    positions,
		EmployeeID:   employeeID,
		ClientID:     clientID,
		Status:       Created,
		DeliveryType: deliveryType,
		Total:        total,
		CreatedAt:    now,
	}
}

func (o *Order) AddPosition(position *Position, now time.Time) {
	for _, orderPosition := range o.Positions {
		if orderPosition.PositionID == position.ID {
			orderPosition.Quantity++
			orderPosition.UnitPrice = position.Price * float64(orderPosition.Quantity)
			o.Total += position.Price
			o.UpdatedAt = now
		} else {
			o.Positions = append(o.Positions, orderPosition)
			o.Total += orderPosition.UnitPrice
		}
	}
}

func (o *Order) ChangeStatus(newStatus OrderStatus, now time.Time) {
	o.Status = newStatus
	o.UpdatedAt = now
}
