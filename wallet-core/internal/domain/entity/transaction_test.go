package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransaction_NewTransaction(t *testing.T) {
	customerFoo, err := NewCustomer("foo", "foo@gmail.com")
	require.NoError(t, err)

	accountFoo := NewAccount(customerFoo)
	accountFoo.Credit(1000)

	customerBar, err := NewCustomer("bar", "bar@gmail.com")
	require.NoError(t, err)

	accountBar := NewAccount(customerBar)
	accountBar.Credit(1000)

	transaction, err := NewTransaction(&accountFoo, &accountBar, 100)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, accountFoo.Balance, float64(900))
	require.Equal(t, accountBar.Balance, float64(1100))

	_, err = NewTransaction(&accountFoo, &accountBar, 1000)
	require.Error(t, err)
}
