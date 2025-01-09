package repository

import "github.com/lucasmdomingues/wallet-core/internal/domain/entity"

type Account interface {
	Save(account entity.Account) error
	FindByID(id string) (entity.Account, error)
	UpdateBalance(account entity.Account) error
}
