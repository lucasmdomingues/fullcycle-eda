package repository

import "github.com/lucasmdomingues/wallet-core/internal/domain/entity"

type Customer interface {
	Get(id string) (entity.Customer, error)
	Save(customer entity.Customer) error
}
