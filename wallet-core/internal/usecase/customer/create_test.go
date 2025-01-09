package customer

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

func TestCreateCustomerUsecase_Execute(t *testing.T) {
	m := &customerRepositoryMock{}
	m.On("Save", mock.Anything).Return(nil)

	usecase := NewCreateCustomerUsecase(m)

	output, err := usecase.Execute(CreateCustomerInputDTO{
		Name:  "John Doe",
		Email: "j@j",
	})
	require.NoError(t, err)
	require.NotEmpty(t, output.ID)
	require.Equal(t, "John Doe", output.Name)
	require.Equal(t, "j@j", output.Email)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
