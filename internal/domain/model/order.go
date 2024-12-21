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
	ID        int64
	Positions []*Position
	CreatedBy int64
	Client    int64
	Status    OrderStatus
	DelType   DeliveryType
	Total     float32
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewOrder(positions []*Position, createdBy, client int64, delType DeliveryType, now time.Time) *Order {
	var total float32
	for _, position := range positions {
		total += position.Price
	}

	return &Order{
		Positions: positions,
		CreatedBy: createdBy,
		Client:    client,
		Status:    Created,
		DelType:   delType,
		Total:     total,
		CreatedAt: now,
	}
}
