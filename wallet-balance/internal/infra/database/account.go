package database

import (
	"database/sql"
	"log"

	"github.com/lucasmdomingues/wallet-balance/internal/domain/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{db}
}

func (a *AccountDB) UpdateBalance(account entity.Account) error {
	stmt, err := a.DB.Prepare("UPDATE accounts SET balance = ? WHERE id = ?")
	if err != nil {
		log.Println("failed to prepare query", err)
		return err
	}

	_, err = stmt.Exec(account.Balance, account.ID)
	if err != nil {
		log.Println("failed to execute query", err)
		return err
	}

	return nil
}

func (a *AccountDB) FindByID(id string) (entity.Account, error) {
	var account entity.Account

	stmt, err := a.DB.Prepare("SELECT * FROM accounts WHERE id = ?")
	if err != nil {
		log.Println("failed to prepare query", err)
		return entity.Account{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&account.ID,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		log.Println("failed to execute query", err)
		return entity.Account{}, err
	}

	return account, nil
}
