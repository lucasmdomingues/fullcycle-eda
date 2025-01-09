package database

import (
	"database/sql"
	"testing"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite

	db         *sql.DB
	repository *AccountDB
}

func (suite *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	suite.db = db

	_, err = db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at DATE)")
	suite.NoError(err)

	_, err = db.Exec("CREATE TABLE accounts (id varchar(255), customer_id varchar(255), balance int, created_at DATE)")
	suite.NoError(err)

	suite.repository = NewAccountDB(db)
}

func (suite *AccountDBTestSuite) TearDownSuite() {
	_, err := suite.db.Exec("DROP TABLE customers")
	suite.NoError(err)

	_, err = suite.db.Exec("DROP TABLE accounts")
	suite.NoError(err)

	err = suite.db.Close()
	suite.NoError(err)
}

func (suite *AccountDBTestSuite) TestSave() {
	customer, err := entity.NewCustomer("John Doe", "j@j.com")
	suite.NoError(err)

	account := entity.NewAccount(customer)

	err = suite.repository.Save(account)
	suite.NoError(err)
}

func (suite *AccountDBTestSuite) TestFindByID() {
	customer, err := entity.NewCustomer("John Doe", "j@j.com")
	suite.NoError(err)

	_, err = suite.db.Exec("INSERT INTO customers (id, name, email, created_at) VALUES (?,?,?,?)",
		customer.ID, customer.Name, customer.Email, customer.CreatedAt)
	suite.NoError(err)

	account := entity.NewAccount(customer)

	err = suite.repository.Save(account)
	suite.NoError(err)

	accountDB, err := suite.repository.FindByID(account.ID)
	suite.NoError(err)

	suite.Equal(accountDB.ID, account.ID)
	suite.Equal(accountDB.Customer.ID, account.Customer.ID)
	suite.Equal(accountDB.Customer.Name, account.Customer.Name)
	suite.Equal(accountDB.Customer.Email, account.Customer.Email)
	suite.Equal(accountDB.Balance, account.Balance)
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}
