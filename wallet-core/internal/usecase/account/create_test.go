package account

import (
	"testing"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type customerRepositoryMock struct {
	mock.Mock
}

func (m *customerRepositoryMock) Get(id string) (entity.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Customer), args.Error(1)
}

func (m *customerRepositoryMock) Save(customer entity.Customer) error {
	args := m.Called(customer)
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

func TestCreateAccountUsecase_Execute(t *testing.T) {
	customer, err := entity.NewCustomer("John Wick", "j@j")
	require.NoError(t, err)

	customerRepositoryMock := &customerRepositoryMock{}
	customerRepositoryMock.On("Get", customer.ID).Return(customer, nil)

	accountRepositoryMock := &accountRepositoryMock{}
	accountRepositoryMock.On("Save", mock.Anything).Return(nil)

	usecase := NewCreateAccountUsecase(accountRepositoryMock, customerRepositoryMock)

	output, err := usecase.Execute(CreateAccountInputDTO{
		CustomerID: customer.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, output)

	customerRepositoryMock.AssertExpectations(t)
	accountRepositoryMock.AssertExpectations(t)

	customerRepositoryMock.AssertNumberOfCalls(t, "Get", 1)
	accountRepositoryMock.AssertNumberOfCalls(t, "Save", 1)
}
