package account

import (
	"log"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
	"github.com/lucasmdomingues/wallet-core/internal/domain/repository"
)

type CreateAccountUsecase struct {
	AccountRepository  repository.Account
	CustomerRepository repository.Customer
}

func NewCreateAccountUsecase(accountRepository repository.Account, customerRepository repository.Customer) *CreateAccountUsecase {
	return &CreateAccountUsecase{
		AccountRepository:  accountRepository,
		CustomerRepository: customerRepository,
	}
}

func (usecase *CreateAccountUsecase) Execute(input CreateAccountInputDTO) (CreateAccountOutputDTO, error) {
	customer, err := usecase.CustomerRepository.Get(input.CustomerID)
	if err != nil {
		log.Println("failed to get customer", err)
		return CreateAccountOutputDTO{}, err
	}

	account := entity.NewAccount(customer)

	err = usecase.AccountRepository.Save(account)
	if err != nil {
		log.Println("failed to save account", err)
		return CreateAccountOutputDTO{}, err
	}

	return CreateAccountOutputDTO{account.ID}, nil
}
