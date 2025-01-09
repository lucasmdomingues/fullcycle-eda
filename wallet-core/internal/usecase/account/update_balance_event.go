package account

import "time"

type BalanceUpdatedEvent struct {
	Name    string
	Payload any
}

func NewBalanceUpdatedEvent() *BalanceUpdatedEvent {
	return &BalanceUpdatedEvent{
		Name: "BalanceUpdated",
	}
}

func (e *BalanceUpdatedEvent) GetName() string {
	return e.Name
}

func (e *BalanceUpdatedEvent) GetPayload() any {
	return e.Payload
}

func (e *BalanceUpdatedEvent) SetPayload(payload any) {
	e.Payload = payload
}

func (e *BalanceUpdatedEvent) GetDateTime() time.Time {
	return time.Now()
}
