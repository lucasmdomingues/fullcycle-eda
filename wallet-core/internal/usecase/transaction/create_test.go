package transaction

import (
	"context"
	"testing"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
	"github.com/lucasmdomingues/wallet-core/internal/usecase/account"
	"github.com/lucasmdomingues/wallet-core/pkg/events"
	"github.com/lucasmdomingues/wallet-core/pkg/uow"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type transactionRepositoryMock struct {
	mock.Mock
}

func (m *transactionRepositoryMock) Create(transaction entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type accountRepositoryMock struct {
	mock.Mock
}

func (m *accountRepositoryMock) Save(account entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *accountRepositoryMock) FindByID(id string) (entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Account), args.Error(1)
}

func (m *accountRepositoryMock) UpdateBalance(account entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateTransactionUsecase_Execute(t *testing.T) {
	uowMock := &uow.UowMock{}
	uowMock.On("Do", mock.Anything, mock.Anything).Return(nil)

	customerFrom, err := entity.NewCustomer("Customer From", "customer_from@mail.com")
	require.NoError(t, err)

	accountFrom := entity.NewAccount(customerFrom)
	accountFrom.Credit(1000)

	customerTo, err := entity.NewCustomer("Customer To", "customer_to@mail.com")
	require.NoError(t, err)

	accountTo := entity.NewAccount(customerTo)
	accountTo.Credit(1000)

	dispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := NewTransactionCreatedEvent()
	balanceUpdatedEvent := account.NewBalanceUpdatedEvent()

	usecase := NewCreateTransactionUsecase(
		uowMock,
		dispatcher,
		transactionCreatedEvent,
		balanceUpdatedEvent,
	)

	input := CreateTransactionInputDTO{
		AccountIDFrom: accountFrom.ID,
		AccountIDTo:   accountTo.ID,
		Amount:        100,
	}

	_, err = usecase.Execute(context.Background(), input)
	require.NoError(t, err)

	uowMock.AssertExpectations(t)
	uowMock.AssertNumberOfCalls(t, "Do", 1)
}
