package balance

import (
	"github.com/lucasmdomingues/wallet-balance/internal/domain/repository"
)

type FindByAccountIDUsecase struct {
	balanceDB repository.Balance
}

func NewFindByAccountIDUsecase(balanceDB repository.Balance) *FindByAccountIDUsecase {
	return &FindByAccountIDUsecase{balanceDB}
}

func (usecase *FindByAccountIDUsecase) Execute(id string) (FindByAccountIDOutput, error) {
	balance, err := usecase.balanceDB.FindByAccountID(id)
	if err != nil {
		return FindByAccountIDOutput{}, err
	}

	return FindByAccountIDOutput{
		ID:        balance.ID,
		AccountID: balance.AccountID,
		Amount:    balance.Amount,
		CreatedAt: balance.CreatedAt,
	}, nil
}
