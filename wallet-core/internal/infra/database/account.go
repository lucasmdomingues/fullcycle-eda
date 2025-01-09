package database

import (
	"database/sql"
	"log"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{db}
}

func (a *AccountDB) FindByID(id string) (entity.Account, error) {
	var account entity.Account

	stmt, err := a.DB.Prepare(`
		SELECT a.id, a.balance, a.created_at, c.id, c.name, c.email, c.created_at
		FROM accounts a
		INNER JOIN customers c ON a.customer_id = c.id
		WHERE a.id = ?
	`)
	if err != nil {
		log.Println("failed to prepare query", err)
		return entity.Account{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&account.ID,
		&account.Balance,
		&account.CreatedAt,
		&account.Customer.ID,
		&account.Customer.Name,
		&account.Customer.Email,
		&account.Customer.CreatedAt,
	)
	if err != nil {
		log.Println("failed to execute query", err)
		return entity.Account{}, err
	}

	return account, nil
}

func (a *AccountDB) Save(account entity.Account) error {
	stmt, err := a.DB.Prepare("INSERT INTO accounts (id, customer_id, balance, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("failed to prepare query", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(account.ID, account.Customer.ID, account.Balance, account.CreatedAt)
	if err != nil {
		log.Println("failed to execute query", err)
		return err
	}

	return nil
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
