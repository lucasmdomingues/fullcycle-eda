package account

import (
	"testing"
	"time"

	"github.com/lucasmdomingues/wallet-balance/internal/domain/entity"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type accountDBMock struct {
	mock.Mock
}

// Save implements repository.Account.
func (a *accountDBMock) FindByID(id string) (entity.Account, error) {
	args := a.Called(id)
	return args.Get(0).(entity.Account), args.Error(1)
}

// UpdateBalance implements repository.Account.
func (a *accountDBMock) UpdateBalance(account entity.Account) error {
	panic("unimplemented")
}

func TestFindByID_Execute(t *testing.T) {
	accountMocked := entity.Account{
		ID:        "1",
		Balance:   100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	accountDB := &accountDBMock{}
	accountDB.On("FindByID", mock.Anything).Return(accountMocked, nil)

	usecase := NewFindByIDUsecase(accountDB)

	output, err := usecase.Execute("1")
	require.NoError(t, err)

	require.Equal(t, accountMocked.ID, output.ID)
	require.Equal(t, accountMocked.Balance, output.Balance)
	require.Equal(t, accountMocked.CreatedAt, output.CreatedAt)
	require.Equal(t, accountMocked.UpdatedAt, output.UpdatedAt)

	accountDB.AssertNumberOfCalls(t, "FindByID", 1)
	accountDB.AssertExpectations(t)
}
