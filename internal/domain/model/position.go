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
	ID        int64
	Name      string
	Barcode   string
	Price     float32
	PosType   PositionType
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewPosition(name, barcode string, price float32, posType PositionType, now time.Time) *Position {
	return &Position{
		Name:      name,
		Barcode:   barcode,
		Price:     price,
		PosType:   posType,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
