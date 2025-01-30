package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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
	Quantity  int64
	UnitPrice float64
	Position  *Position
}

type Order struct {
	ID           int64
	Positions    map[int64]*OrderPosition
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
		Positions:    make(map[int64]*OrderPosition),
		ClientID:     clientID,
		Status:       Created,
		DeliveryType: deliveryType,
		CreatedAt:    now,
	}
}

func (op *OrderPosition) Value() (driver.Value, error) {
	if op.Position == nil {

		return "[]", nil
	}

	positionsJSON, err := json.Marshal(op.Position)
	if err != nil {

		return "", fmt.Errorf("failed to marshal positions: %w", err)
	}

	return string(positionsJSON), nil
}

func (op *OrderPosition) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return errors.New("not a string")
	}

	return json.Unmarshal([]byte(b), &op)
}

func (o *Order) ChangeStatus(newStatus OrderStatus, now time.Time) {
	o.Status = newStatus
	o.UpdatedAt = now
}

func (o *Order) AddPositions(positions []*Position) {
	for _, position := range positions {
		orderPosition, has := o.Positions[position.ID]
		if has {
			orderPosition.Quantity += 1
			orderPosition.UnitPrice = float64(orderPosition.Quantity) * position.Price
			o.Positions[position.ID] = orderPosition
		} else {
			o.Positions[position.ID] = &OrderPosition{
				Position:  position,
				Quantity:  1,
				UnitPrice: position.Price,
			}
		}
	}
	for _, orderPosition := range o.Positions {
		o.Total += orderPosition.Position.Price
	}
}
