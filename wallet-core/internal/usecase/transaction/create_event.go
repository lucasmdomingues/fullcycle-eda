package transaction

import "time"

type TransactionCreatedEvent struct {
	Name    string
	Payload any
}

func NewTransactionCreatedEvent() *TransactionCreatedEvent {
	return &TransactionCreatedEvent{
		Name: "TransactionCreated",
	}
}

func (e *TransactionCreatedEvent) GetName() string {
	return e.Name
}

func (e *TransactionCreatedEvent) GetPayload() any {
	return e.Payload
}

func (e *TransactionCreatedEvent) SetPayload(payload any) {
	e.Payload = payload
}

func (e *TransactionCreatedEvent) GetDateTime() time.Time {
	return time.Now()
}
