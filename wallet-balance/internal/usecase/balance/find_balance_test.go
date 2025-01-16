package balance

import (
	"testing"
	"time"

	"github.com/lucasmdomingues/wallet-balance/internal/domain/entity"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type balanceDBMock struct {
	mock.Mock
}

func (a *balanceDBMock) FindByAccountID(id string) (entity.Balance, error) {
	args := a.Called(id)
	return args.Get(0).(entity.Balance), args.Error(1)
}

func (a *balanceDBMock) SaveBalance(account entity.Balance) error {
	panic("unimplemented")
}

func TestFindByID_Execute(t *testing.T) {
	balanceMocked := entity.Balance{
		ID:        1,
		AccountID: "1",
		Amount:    100,
		CreatedAt: time.Now(),
	}

	balanceDB := &balanceDBMock{}
	balanceDB.On("FindByAccountID", mock.Anything).Return(balanceMocked, nil)

	usecase := NewFindByAccountIDUsecase(balanceDB)

	output, err := usecase.Execute("1")
	require.NoError(t, err)

	require.Equal(t, balanceMocked.ID, output.ID)
	require.Equal(t, balanceMocked.AccountID, output.AccountID)
	require.Equal(t, balanceMocked.Amount, output.Amount)
	require.Equal(t, balanceMocked.CreatedAt, output.CreatedAt)

	balanceDB.AssertNumberOfCalls(t, "FindByAccountID", 1)
	balanceDB.AssertExpectations(t)
}
