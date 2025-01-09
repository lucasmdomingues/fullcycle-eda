package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccount_NewAccount(t *testing.T) {
	customer, err := NewCustomer("John Doe", "j@j")
	require.NoError(t, err)

	account := NewAccount(customer)
	require.Equal(t, account.Customer.ID, customer.ID)
	require.Equal(t, account.Customer.Name, customer.Name)
	require.Equal(t, account.Customer.Email, customer.Email)
}

func TestAccount_Credit(t *testing.T) {
	customer, err := NewCustomer("John Doe", "j@j")
	require.NoError(t, err)

	account := NewAccount(customer)
	account.Credit(10)
	require.Equal(t, account.Balance, float64(10))
}

func TestAccount_Debit(t *testing.T) {
	customer, err := NewCustomer("John Doe", "j@j")
	require.NoError(t, err)

	account := NewAccount(customer)
	account.Credit(10)
	account.Debit(10)
	require.Equal(t, account.Balance, float64(0))
}
