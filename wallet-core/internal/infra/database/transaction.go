package database

import (
	"database/sql"
	"log"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{db}
}

func (t *TransactionDB) Create(transaction entity.Transaction) error {
	stmt, err := t.DB.Prepare(`INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) VALUES (?,?,?,?,?)`)
	if err != nil {
		log.Println("failed to prepare query", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountTo.ID,
		transaction.Amount, transaction.CreatedAt)
	if err != nil {
		log.Println("failed to execute query", err)
		return err
	}

	return nil
}
