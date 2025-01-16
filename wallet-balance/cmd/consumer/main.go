package main

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	qm "github.com/lucasmdomingues/wallet-balance/internal/infra/queue-manager"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("start consumer...")

	balancesConsumer := qm.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}, []string{"wallet.core.balances"})

	balancesMsgChan := make(chan *kafka.Message)
	defer close(balancesMsgChan)

	go func() {
		err := balancesConsumer.Consume(balancesMsgChan)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	for {
		msg := <-balancesMsgChan
		fmt.Println("message received", string(msg.Value))
	}
}
