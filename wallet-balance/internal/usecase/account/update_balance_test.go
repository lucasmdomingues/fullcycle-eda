package account

import (
	"context"
	"testing"

	"github.com/lucasmdomingues/wallet-balance/pkg/uow"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateBalanceUsecase_Execute(t *testing.T) {
	uowMock := &uow.UowMock{}
	uowMock.On("Do", mock.Anything, mock.Anything).Return(nil)

	usecase := NewUpdateBalanceUsecase(uowMock)
	ctx := context.Background()

	input := UpdateBalanceInputDTO{
		AccountIDFrom:        "5",
		AccountIDTo:          "3",
		BalanceAccountIDFrom: 10,
		BalanceAccountIDTo:   20,
	}

	err := usecase.UpdateBalance(ctx, input)
	require.NoError(t, err)

	uowMock.AssertExpectations(t)
	uowMock.AssertNumberOfCalls(t, "Do", 1)
}
