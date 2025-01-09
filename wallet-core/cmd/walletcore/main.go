package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lucasmdomingues/wallet-core/internal/infra/api"
	"github.com/lucasmdomingues/wallet-core/internal/infra/api/handler"
	"github.com/lucasmdomingues/wallet-core/internal/infra/database"
	qm "github.com/lucasmdomingues/wallet-core/internal/infra/queue-manager"
	"github.com/lucasmdomingues/wallet-core/internal/usecase/account"
	"github.com/lucasmdomingues/wallet-core/internal/usecase/customer"
	"github.com/lucasmdomingues/wallet-core/internal/usecase/transaction"
	"github.com/lucasmdomingues/wallet-core/pkg/events"
	"github.com/lucasmdomingues/wallet-core/pkg/uow"
)

func main() {
	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	defer db.Close()

	customerDB := database.NewCustomerDB(db)
	accountDB := database.NewAccountDB(db)
	transactionDB := database.NewTransactionDB(db)

	uow := uow.NewUow(ctx, db)
	uow.Register("AccountRepository", func(tx *sql.Tx) interface{} {
		return accountDB
	})
	uow.Register("TransactionRepository", func(tx *sql.Tx) interface{} {
		return transactionDB
	})

	qmProducer := qm.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	})

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", transaction.NewTransactionCreatedHandler(qmProducer))
	eventDispatcher.Register("BalanceUpdated", account.NewBalanceUpdatedEventHandler(qmProducer))

	createTransactionEvent := transaction.NewTransactionCreatedEvent()
	updateBalanceEvent := account.NewBalanceUpdatedEvent()

	createCustomerUsecase := customer.NewCreateCustomerUsecase(customerDB)
	createAccountUsecase := account.NewCreateAccountUsecase(accountDB, customerDB)
	createTransactionUsecase := transaction.NewCreateTransactionUsecase(
		uow,
		eventDispatcher,
		createTransactionEvent,
		updateBalanceEvent,
	)

	customerHandler := handler.NewCustomerHandler(createCustomerUsecase)
	accountHandler := handler.NewAccountHandler(createAccountUsecase)
	transactionHandler := handler.NewTransactionHandler(createTransactionUsecase)

	server := api.NewServer(":8080")
	server.AddRoute("/customers", customerHandler.CreateCustomer)
	server.AddRoute("/accounts", accountHandler.CreateAccount)
	server.AddRoute("/transactions", transactionHandler.CreateTransaction)

	server.Start()
}
