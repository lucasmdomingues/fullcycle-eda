package balance

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

	usecase := NewSaveBalanceUsecase(uowMock)
	ctx := context.Background()

	input := SaveBalanceInputDTO{
		AccountIDFrom:        "5",
		AccountIDTo:          "3",
		BalanceAccountIDFrom: 10,
		BalanceAccountIDTo:   20,
	}

	err := usecase.Execute(ctx, input)
	require.NoError(t, err)

	uowMock.AssertExpectations(t)
	uowMock.AssertNumberOfCalls(t, "Do", 1)
}
