package account

import (
	"log"

	qm "github.com/lucasmdomingues/wallet-core/internal/infra/queue-manager"
	"github.com/lucasmdomingues/wallet-core/pkg/events"
)

type BalanceUpdatedEventHandler struct {
	Producer *qm.Producer
}

func NewBalanceUpdatedEventHandler(producer *qm.Producer) *BalanceUpdatedEventHandler {
	return &BalanceUpdatedEventHandler{producer}
}

func (h *BalanceUpdatedEventHandler) Handle(message events.Event) {
	err := h.Producer.Publish(message.GetPayload(), nil, "wallet.core.balances")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Message published", message.GetName())
}
