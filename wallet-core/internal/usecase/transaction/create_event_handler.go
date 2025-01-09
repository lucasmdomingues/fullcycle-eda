package transaction

import (
	"log"

	qm "github.com/lucasmdomingues/wallet-core/internal/infra/queue-manager"
	"github.com/lucasmdomingues/wallet-core/pkg/events"
)

type TransactionCreatedEventHandler struct {
	Producer *qm.Producer
}

func NewTransactionCreatedHandler(producer *qm.Producer) *TransactionCreatedEventHandler {
	return &TransactionCreatedEventHandler{producer}
}

func (h *TransactionCreatedEventHandler) Handle(message events.Event) {
	err := h.Producer.Publish(message.GetPayload(), nil, "wallet.core.transactions")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Message published", message.GetName())
}
