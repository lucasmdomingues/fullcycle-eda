package repository

import "github.com/lucasmdomingues/wallet-balance/internal/domain/entity"

type Account interface {
	FindByID(id string) (entity.Account, error)
	UpdateBalance(account entity.Account) error
}
