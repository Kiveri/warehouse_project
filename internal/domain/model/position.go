package model

import "time"

type PositionType int

const (
	BasicProduct    PositionType = 1
	BasicConsumable PositionType = 2
	Liquid          PositionType = 3
	OversizeProduct PositionType = 4
)

type Position struct {
	ID           int64
	Name         string
	Barcode      string
	Price        float32
	PositionType PositionType
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func NewPosition(name, barcode string, price float32, positionType PositionType, now time.Time) *Position {
	return &Position{
		Name:         name,
		Barcode:      barcode,
		Price:        price,
		PositionType: positionType,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (p *Position) ChangePrice(newPrice float32, now time.Time) {
	p.Price = newPrice
	p.UpdatedAt = now
}
