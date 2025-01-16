package entity

import (
	"time"
)

type Balance struct {
	ID        int
	AccountID string
	Amount    float64
	CreatedAt time.Time
}
