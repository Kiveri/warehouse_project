package model

import "time"

type Client struct {
	ID          int64
	Name        string
	Phone       string
	Email       string
	HomeAddress string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewClient(name, phone, email, homeAddress string, now time.Time) *Client {
	return &Client{
		Name:        name,
		Phone:       phone,
		Email:       email,
		HomeAddress: homeAddress,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (c *Client) ChangeAddress(newAddress string) {
	c.HomeAddress = newAddress
	c.UpdatedAt = time.Now()
}
