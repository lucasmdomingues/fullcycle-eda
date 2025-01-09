package database

import (
	"database/sql"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite

	db         *sql.DB
	repository *TransactionDB
}

func (suite *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	suite.db = db

	_, err = db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at DATE)")
	suite.NoError(err)

	_, err = db.Exec("CREATE TABLE accounts (id varchar(255), customer_id varchar(255), balance int, created_at DATE)")
	suite.NoError(err)

	_, err = db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")
	suite.NoError(err)

	suite.repository = NewTransactionDB(db)
}

func (suite *TransactionDBTestSuite) TearDownSuite() {
	_, err := suite.db.Exec("DROP TABLE customers")
	suite.NoError(err)

	_, err = suite.db.Exec("DROP TABLE accounts")
	suite.NoError(err)

	_, err = suite.db.Exec("DROP TABLE transactions")
	suite.NoError(err)

	err = suite.db.Close()
	suite.NoError(err)
}

func (suite *TransactionDBTestSuite) TestCreate() {
	customerFrom, err := entity.NewCustomer("Customer from", "from@mail.com")
	suite.NoError(err)

	accountFrom := entity.NewAccount(customerFrom)
	accountFrom.Credit(1000)

	customerTo, err := entity.NewCustomer("Customer to", "to@mail.com")
	suite.NoError(err)

	accountTo := entity.NewAccount(customerTo)
	accountTo.Credit(1000)

	transaction, err := entity.NewTransaction(&accountFrom, &accountTo, 100.00)
	suite.NoError(err)

	err = suite.repository.Create(transaction)
	suite.NoError(err)
}
