package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

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
	Price        float64
	PositionType PositionType
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewPosition(name, barcode string, price float64, positionType PositionType, now time.Time) *Position {
	return &Position{
		Name:         name,
		Barcode:      barcode,
		Price:        price,
		PositionType: positionType,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (p *Position) Value() (driver.Value, error) {
	positionJSON, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("failed to marshal position: %w", err)
	}

	return string(positionJSON), nil
}

func (p *Position) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return errors.New("not a string")
	}

	return json.Unmarshal([]byte(b), &p)
}

func (p *Position) ChangePrice(newPrice float64, now time.Time) {
	p.Price = newPrice
	p.UpdatedAt = now
}
