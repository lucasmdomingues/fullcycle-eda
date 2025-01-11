package account

import (
	"context"
	"log"

	"github.com/lucasmdomingues/wallet-balance/internal/domain/entity"
	"github.com/lucasmdomingues/wallet-balance/internal/domain/repository"
	"github.com/lucasmdomingues/wallet-balance/pkg/uow"
)

type UpdateBalanceUsecase struct {
	uow uow.UowInterface
}

func NewUpdateBalanceUsecase(uow uow.UowInterface) *UpdateBalanceUsecase {
	return &UpdateBalanceUsecase{uow}
}

func (usecase *UpdateBalanceUsecase) UpdateBalance(ctx context.Context, input UpdateBalanceInputDTO) error {
	err := usecase.uow.Do(ctx, func(uow *uow.Uow) error {
		accountRepository, err := getAccountRepository(ctx, uow)
		if err != nil {
			return err
		}

		accountFrom := entity.Account{
			ID:      input.AccountIDFrom,
			Balance: input.BalanceAccountIDFrom,
		}

		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			log.Println("failed to update account from balance", err)
			return err
		}

		accountTo := entity.Account{
			ID:      input.AccountIDTo,
			Balance: input.BalanceAccountIDTo,
		}

		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			log.Println("failed to update account to balance", err)
			return err
		}

		return nil
	})

	return err
}

func getAccountRepository(ctx context.Context, uow uow.UowInterface) (repository.Account, error) {
	accountRepository, err := uow.GetRepository(ctx, "AccountRepository")
	if err != nil {
		log.Println("failed to get account repository", err)
		return nil, err
	}
	return accountRepository.(repository.Account), nil
}
