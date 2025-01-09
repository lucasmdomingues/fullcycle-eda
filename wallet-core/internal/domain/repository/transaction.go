package repository

import "github.com/lucasmdomingues/wallet-core/internal/domain/entity"

type Transaction interface {
	Create(transaction entity.Transaction) error
}
