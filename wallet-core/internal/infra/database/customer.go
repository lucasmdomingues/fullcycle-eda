package database

import (
	"database/sql"
	"log"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
)

type CustomerDB struct {
	DB *sql.DB
}

func NewCustomerDB(db *sql.DB) *CustomerDB {
	return &CustomerDB{db}
}

func (r *CustomerDB) Get(id string) (entity.Customer, error) {
	stmt, err := r.DB.Prepare("SELECT id, name, email, created_at FROM customers WHERE id = ?")
	if err != nil {
		log.Println("failed to prepare query", err)
		return entity.Customer{}, err
	}
	defer stmt.Close()

	var customer entity.Customer

	err = stmt.QueryRow(id).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.CreatedAt)
	if err != nil {
		log.Println("failed to execute query", err)
		return entity.Customer{}, err
	}

	return customer, nil
}

func (r *CustomerDB) Save(customer entity.Customer) error {
	stmt, err := r.DB.Prepare("INSERT INTO customers (id, name, email, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("failed to prepare query", err)
		return err
	}

	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email, customer.CreatedAt)
	if err != nil {
		log.Println("failed to execute query", err)
		return err
	}

	return nil
}
