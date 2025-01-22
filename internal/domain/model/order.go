package model

import (
	"encoding/json"
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
	PositionID int64
	Quantity   int64
	UnitPrice  float64
	Position   *Position
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

func (o *Order) Value() (string, error) {
	if o.Positions == nil {

		return "", nil
	}

	positionsJSON, err := json.Marshal(o.Positions)
	if err != nil {

		return "", fmt.Errorf("failed to marshal positions: %w", err)
	}

	return string(positionsJSON), nil
}

/*func (op *OrderPosition) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return errors.New("not a string")
	}

	return json.Unmarshal([]byte(b), &op)
}

*/

func (o *Order) ChangeStatus(newStatus OrderStatus, now time.Time) {
	o.Status = newStatus
	o.UpdatedAt = now
}

func (o *Order) AddPositions(positionID int64, quantity int64, unitPrice float64) error {
	for _, orderPosition := range o.Positions {
		if orderPosition.PositionID == positionID {
			orderPosition.Quantity += quantity
			orderPosition.UnitPrice = float64(orderPosition.Quantity) * unitPrice

			return nil
		}
	}

	newOrderPosition := &OrderPosition{
		PositionID: positionID,
		Quantity:   quantity,
		UnitPrice:  float64(quantity) * unitPrice,
	}
	o.Positions = append(o.Positions, newOrderPosition)

	return nil
}
