package account

type CreateAccountInputDTO struct {
	CustomerID string `json:"customer_id"`
}

type CreateAccountOutputDTO struct {
	ID string `json:"id"`
}

type UpdateBalanceOutputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}
