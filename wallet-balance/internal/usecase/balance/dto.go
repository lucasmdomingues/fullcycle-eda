package balance

import "time"

type SaveBalanceInputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

type FindByAccountIDOutput struct {
	ID        int       `json:"id"`
	AccountID string    `json:"account_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
