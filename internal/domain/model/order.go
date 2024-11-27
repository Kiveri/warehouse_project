package model

import "time"

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

type Order struct {
	ID        int
	Positions []Position
	Status    OrderStatus
	DelType   DeliveryType
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewOrder(positions []Position, status OrderStatus, delType DeliveryType, now time.Time) *Order {
	return &Order{
		Positions: positions,
		Status:    status,
		DelType:   delType,
		CreatedAt: now,
	}
}
