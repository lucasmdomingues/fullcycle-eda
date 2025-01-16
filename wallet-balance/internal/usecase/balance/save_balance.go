package balance

import (
	"context"
	"log"

	"github.com/lucasmdomingues/wallet-balance/internal/domain/entity"
	"github.com/lucasmdomingues/wallet-balance/internal/domain/repository"
	"github.com/lucasmdomingues/wallet-balance/pkg/uow"
)

type SaveBalanceUsecase struct {
	uow uow.UowInterface
}

func NewSaveBalanceUsecase(uow uow.UowInterface) *SaveBalanceUsecase {
	return &SaveBalanceUsecase{uow}
}

func (usecase *SaveBalanceUsecase) SaveBalance(ctx context.Context, input SaveBalanceInputDTO) error {
	err := usecase.uow.Do(ctx, func(uow *uow.Uow) error {
		accountRepository, err := getBalanceRepository(ctx, uow)
		if err != nil {
			return err
		}

		accountFrom := entity.Balance{
			AccountID: input.AccountIDFrom,
			Amount:    input.BalanceAccountIDFrom,
		}

		err = accountRepository.SaveBalance(accountFrom)
		if err != nil {
			log.Println("failed to update account from balance", err)
			return err
		}

		accountTo := entity.Balance{
			AccountID: input.AccountIDTo,
			Amount:    input.BalanceAccountIDTo,
		}

		err = accountRepository.SaveBalance(accountTo)
		if err != nil {
			log.Println("failed to update account to balance", err)
			return err
		}

		return nil
	})

	return err
}

func getBalanceRepository(ctx context.Context, uow uow.UowInterface) (repository.Balance, error) {
	balanceRepository, err := uow.GetRepository(ctx, "BalanceRepository")
	if err != nil {
		log.Println("failed to get account repository", err)
		return nil, err
	}
	return balanceRepository.(repository.Balance), nil
}
