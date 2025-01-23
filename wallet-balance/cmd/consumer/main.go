package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lucasmdomingues/wallet-balance/internal/infra/database"
	qm "github.com/lucasmdomingues/wallet-balance/internal/infra/queue-manager"
	"github.com/lucasmdomingues/wallet-balance/internal/usecase/balance"
	"github.com/lucasmdomingues/wallet-balance/pkg/uow"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("connect database...")
	db, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	defer db.Close()

	log.Println("start consumer...")
	balancesConsumer := qm.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}, []string{"wallet.core.balances"})

	balancesMsgChan := make(chan *kafka.Message)
	defer close(balancesMsgChan)

	accountDB := database.NewBalanceDB(db)

	uow := uow.NewUow(ctx, db)
	uow.Register("BalanceRepository", func(tx *sql.Tx) interface{} {
		return accountDB
	})

	saveBalanceUsecase := balance.NewSaveBalanceUsecase(uow)

	log.Println("listening messages...")
	go func() {
		err := balancesConsumer.Consume(balancesMsgChan)
		if err != nil {
			log.Println("failed to consumer msg", err)
			return
		}
	}()

	for {
		msg := <-balancesMsgChan
		log.Println("message received", string(msg.Value))

		var saveBalanceInputDTO balance.SaveBalanceInputDTO

		err := json.Unmarshal(msg.Value, &saveBalanceInputDTO)
		if err != nil {
			log.Println("failed to unmarshal message", err)
			return
		}
		log.Println("message parsed", saveBalanceInputDTO)

		err = saveBalanceUsecase.Execute(ctx, saveBalanceInputDTO)
		if err != nil {
			log.Println("failed to save balance", err)
			return
		}
		log.Println("save balance successfully")
	}
}
