package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/lucasmdomingues/wallet-balance/internal/infra/api"
	"github.com/lucasmdomingues/wallet-balance/internal/infra/database"
	"github.com/lucasmdomingues/wallet-balance/internal/usecase/account"
	"github.com/lucasmdomingues/wallet-balance/pkg/uow"
)

func main() {
	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	defer db.Close()

	accountDB := database.NewAccountDB(db)

	uow := uow.NewUow(ctx, db)
	uow.Register("AccountRepository", func(tx *sql.Tx) interface{} {
		return accountDB
	})

	// balancesConsumer := qm.NewConsumer(&kafka.ConfigMap{
	// 	"bootstrap.servers": "kafka:29092",
	// 	"group.id":          "wallet",
	// }, []string{"wallet.core.balances"})

	// balancesMsgChan := make(chan *kafka.Message)
	// go func() {
	// 	err = balancesConsumer.Consume(balancesMsgChan)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	findAccountByIDUsecase := account.NewFindByIDUsecase(accountDB)

	server := api.NewAPI(findAccountByIDUsecase)

	log.Println("start server...")
	err = server.Start()
	log.Fatal("shutting down server...", err)
}
