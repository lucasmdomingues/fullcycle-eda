package events

import "time"

type Event interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() any
	SetPayload(data any)
}
