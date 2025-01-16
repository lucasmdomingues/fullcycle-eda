package repository

import "github.com/lucasmdomingues/wallet-balance/internal/domain/entity"

type Balance interface {
	FindByAccountID(id string) (entity.Balance, error)
	SaveBalance(balace entity.Balance) error
}
