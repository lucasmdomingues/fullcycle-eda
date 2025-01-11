package database

import (
	"database/sql"
	"testing"
	"time"

	"github.com/lucasmdomingues/wallet-balance/internal/domain/entity"
	_ "github.com/mattn/go-sqlite3"
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

	_, err = db.Exec("CREATE TABLE accounts (id varchar(255), balance int, created_at DATE, updated_at DATE)")
	suite.NoError(err)

	suite.repository = NewAccountDB(db)
}

func (suite *AccountDBTestSuite) TearDownSuite() {
	_, err := suite.db.Exec("DROP TABLE accounts")
	suite.NoError(err)

	err = suite.db.Close()
	suite.NoError(err)
}

func (suite *AccountDBTestSuite) TestFindByID() {
	account := entity.Account{
		ID:        "1",
		Balance:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := suite.db.Exec("INSERT INTO accounts (id, balance, created_at, updated_at) VALUES (?,?,?,?)",
		account.ID, account.Balance, account.CreatedAt, account.UpdatedAt)
	suite.NoError(err)

	account.Balance = 100

	err = suite.repository.UpdateBalance(account)
	suite.NoError(err)

	accountDB, err := suite.repository.FindByID(account.ID)
	suite.NoError(err)

	suite.Equal(accountDB.ID, account.ID)
	suite.Equal(accountDB.Balance, account.Balance)
}

func (suite *AccountDBTestSuite) TestUpdateBalance() {
	account := entity.Account{
		ID:        "1",
		Balance:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := suite.repository.UpdateBalance(account)
	suite.NoError(err)
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}
