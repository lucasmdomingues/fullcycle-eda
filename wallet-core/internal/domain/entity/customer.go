package entity

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string
	Name      string
	Email     string
	Accounts  []Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name string, email string) (Customer, error) {
	customer := Customer{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := customer.Validate()
	if err != nil {
		log.Println("failed to validate customer", err)
		return Customer{}, err
	}

	return customer, nil
}

func (c Customer) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func (c *Customer) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()

	err := c.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) AddAccount(account Account) error {
	if account.Customer.ID != c.ID {
		return errors.New("account does not belong to customer")
	}

	c.Accounts = append(c.Accounts, account)
	return nil
}
