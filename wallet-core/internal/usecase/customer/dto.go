package customer

import "time"

type CreateCustomerInputDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateCustomerOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
