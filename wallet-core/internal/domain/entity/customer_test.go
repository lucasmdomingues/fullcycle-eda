package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomer_NewCustomer(t *testing.T) {
	customer, err := NewCustomer("foo bar", "foo@bar")
	require.NoError(t, err)
	require.Equal(t, customer.Name, "foo bar")
	require.Equal(t, customer.Email, "foo@bar")

	_, err = NewCustomer("", "")
	require.Error(t, err)
}

func TestCustomer_Update(t *testing.T) {
	customer, err := NewCustomer("some name", "some_name@email.com")
	require.NoError(t, err)

	err = customer.Update("edited some name", "somename@email.com")
	require.NoError(t, err)

	err = customer.Update("", "somename@email.com")
	require.Error(t, err)
}

func TestCustomer_AddAccount(t *testing.T) {
	customer, err := NewCustomer("some name", "some_name@email.com")
	require.NoError(t, err)

	account := NewAccount(customer)

	err = customer.AddAccount(account)
	require.NoError(t, err)
	require.Len(t, customer.Accounts, 1)
}
