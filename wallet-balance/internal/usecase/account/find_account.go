package account

import "github.com/lucasmdomingues/wallet-balance/internal/domain/repository"

type FindByIDUsecase struct {
	accountDB repository.Account
}

func NewFindByIDUsecase(accountDB repository.Account) *FindByIDUsecase {
	return &FindByIDUsecase{accountDB}
}

func (usecase *FindByIDUsecase) Execute(id string) (FindByIDOutput, error) {
	account, err := usecase.accountDB.FindByID(id)
	if err != nil {
		return FindByIDOutput{}, err
	}

	return FindByIDOutput{
		ID:        account.ID,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}, nil
}
