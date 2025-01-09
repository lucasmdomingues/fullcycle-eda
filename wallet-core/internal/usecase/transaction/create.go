package transaction

import (
	"context"
	"log"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
	"github.com/lucasmdomingues/wallet-core/internal/domain/repository"
	"github.com/lucasmdomingues/wallet-core/internal/usecase/account"
	"github.com/lucasmdomingues/wallet-core/pkg/events"
	"github.com/lucasmdomingues/wallet-core/pkg/uow"
)

type CreateTransactionUsecase struct {
	Uow                     uow.UowInterface
	eventDispatcher         events.EventDispatcher
	transactionCreatedEvent events.Event
	balanceUpdatedEvent     events.Event
}

func NewCreateTransactionUsecase(
	uow uow.UowInterface,
	eventDispatcher events.EventDispatcher,
	transactionCreatedEvent events.Event,
	balanceUpdatedEvent events.Event,
) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{
		Uow:                     uow,
		eventDispatcher:         eventDispatcher,
		transactionCreatedEvent: transactionCreatedEvent,
		balanceUpdatedEvent:     balanceUpdatedEvent,
	}
}

func (usecase *CreateTransactionUsecase) Execute(ctx context.Context, input CreateTransactionInputDTO) (CreateTransactionOutputDTO, error) {
	var (
		createTransactionOutput CreateTransactionOutputDTO
		updateBalanceOutput     account.UpdateBalanceOutputDTO
	)

	err := usecase.Uow.Do(ctx, func(_ *uow.Uow) error {
		accountRepository := usecase.getAccountRepository(ctx)
		transactionRepository := usecase.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindByID(input.AccountIDFrom)
		if err != nil {
			log.Println("failed to find account from", err)
			return err
		}

		accountTo, err := accountRepository.FindByID(input.AccountIDTo)
		if err != nil {
			log.Println("failed to find account to", err)
			return err
		}

		transaction, err := entity.NewTransaction(&accountFrom, &accountTo, input.Amount)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			log.Println("failed to update account from balance", err)
			return err
		}

		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			log.Println("failed to update account to balance", err)
			return err
		}

		err = transactionRepository.Create(transaction)
		if err != nil {
			log.Println("failed to create transaction", err)
			return err
		}

		createTransactionOutput.ID = transaction.ID
		createTransactionOutput.AccountIDFrom = input.AccountIDFrom
		createTransactionOutput.AccountIDTo = input.AccountIDTo
		createTransactionOutput.Amount = input.Amount

		updateBalanceOutput.AccountIDFrom = input.AccountIDFrom
		updateBalanceOutput.AccountIDTo = input.AccountIDTo
		updateBalanceOutput.BalanceAccountIDFrom = accountFrom.Balance
		updateBalanceOutput.BalanceAccountIDTo = accountTo.Balance

		return nil
	})
	if err != nil {
		return createTransactionOutput, err
	}

	usecase.transactionCreatedEvent.SetPayload(createTransactionOutput)
	err = usecase.eventDispatcher.Dispatch(usecase.transactionCreatedEvent)
	if err != nil {
		return CreateTransactionOutputDTO{}, err
	}

	usecase.balanceUpdatedEvent.SetPayload(updateBalanceOutput)
	err = usecase.eventDispatcher.Dispatch(usecase.balanceUpdatedEvent)
	if err != nil {
		return CreateTransactionOutputDTO{}, err
	}

	return createTransactionOutput, nil
}

func (usecase CreateTransactionUsecase) getAccountRepository(ctx context.Context) repository.Account {
	repo, err := usecase.Uow.GetRepository(ctx, "AccountRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.Account)
}

func (usecase CreateTransactionUsecase) getTransactionRepository(ctx context.Context) repository.Transaction {
	repo, err := usecase.Uow.GetRepository(ctx, "TransactionRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.Transaction)
}
