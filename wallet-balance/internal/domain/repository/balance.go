package repository

import "github.com/lucasmdomingues/wallet-balance/internal/domain/entity"

type Account interface {
	Save(account entity.Account) error
	UpdateBalance(account entity.Account) error
}
