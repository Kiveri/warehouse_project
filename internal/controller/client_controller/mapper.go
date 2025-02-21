package client_controller

import (
	"time"

	"warehouse_project/internal/domain/model"
)

type clientResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	HomeAddress string    `json:"home_address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func mapClientToResponse(client *model.Client) *clientResponse {
	return &clientResponse{
		ID:          client.ID,
		Name:        client.Name,
		Phone:       client.Phone,
		Email:       client.Email,
		HomeAddress: client.HomeAddress,
		CreatedAt:   client.CreatedAt,
		UpdatedAt:   client.UpdatedAt,
	}
}
