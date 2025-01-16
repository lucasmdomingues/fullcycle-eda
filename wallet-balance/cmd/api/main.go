package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/lucasmdomingues/wallet-balance/internal/infra/api"
	"github.com/lucasmdomingues/wallet-balance/internal/infra/database"
	"github.com/lucasmdomingues/wallet-balance/internal/usecase/balance"

	"github.com/lucasmdomingues/wallet-balance/pkg/uow"
)

func main() {
	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	defer db.Close()

	accountDB := database.NewBalanceDB(db)

	uow := uow.NewUow(ctx, db)
	uow.Register("BalanceRepository", func(tx *sql.Tx) interface{} {
		return accountDB
	})

	findAccountByIDUsecase := balance.NewFindByAccountIDUsecase(accountDB)

	server := api.NewAPI(findAccountByIDUsecase)

	log.Println("start server...")
	err = server.Start()
	log.Fatal("shutting down server...", err)
}
