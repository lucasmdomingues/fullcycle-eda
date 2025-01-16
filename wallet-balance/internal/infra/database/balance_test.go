package database

import (
	"database/sql"
	"testing"
	"time"

	"github.com/lucasmdomingues/wallet-balance/internal/domain/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type BalanceDBTestSuite struct {
	suite.Suite

	db         *sql.DB
	repository *BalanceDB
}

func (suite *BalanceDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	suite.db = db

	_, err = db.Exec("CREATE TABLE balances (id INTEGER PRIMARY KEY AUTOINCREMENT, account_id VARCHAR(255), amount INT, created_at DATE DEFAULT CURRENT_DATE);")
	suite.NoError(err)

	suite.repository = NewBalanceDB(db)
}

func (suite *BalanceDBTestSuite) TearDownSuite() {
	_, err := suite.db.Exec("DROP TABLE balances")
	suite.NoError(err)

	err = suite.db.Close()
	suite.NoError(err)
}

func (suite *BalanceDBTestSuite) TestFindByAccountID() {
	balance := entity.Balance{
		Amount:    100,
		AccountID: "1",
	}

	_, err := suite.db.Exec("INSERT INTO balances (account_id, amount) VALUES (?,?)", balance.AccountID, balance.Amount)
	suite.NoError(err)

	balanceDB, err := suite.repository.FindByAccountID(balance.AccountID)
	suite.NoError(err)

	suite.Equal(balanceDB.AccountID, balance.AccountID)
	suite.Equal(balanceDB.Amount, balance.Amount)
}

func (suite *BalanceDBTestSuite) TestSaveBalances() {
	balance := entity.Balance{
		ID:        1,
		Amount:    100,
		AccountID: "1",
		CreatedAt: time.Now(),
	}

	err := suite.repository.SaveBalance(balance)
	suite.NoError(err)
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}
