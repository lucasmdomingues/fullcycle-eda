package database

import (
	"database/sql"
	"log"

	"github.com/lucasmdomingues/wallet-balance/internal/domain/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{db}
}

func (a *BalanceDB) SaveBalance(balance entity.Balance) error {
	stmt, err := a.DB.Prepare("INSERT INTO balances (account_id, amount) VALUES (?,?)")
	if err != nil {
		log.Println("failed to prepare query", err)
		return err
	}

	_, err = stmt.Exec(balance.AccountID, balance.Amount)
	if err != nil {
		log.Println("failed to execute query", err)
		return err
	}

	return nil
}

func (a *BalanceDB) FindByAccountID(id string) (entity.Balance, error) {
	var balance entity.Balance

	stmt, err := a.DB.Prepare("SELECT * FROM balances WHERE account_id = ? ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		log.Println("failed to prepare query", err)
		return entity.Balance{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&balance.ID,
		&balance.AccountID,
		&balance.Amount,
		&balance.CreatedAt,
	)
	if err != nil {
		log.Println("failed to execute query", err)
		return entity.Balance{}, err
	}

	return balance, nil
}
