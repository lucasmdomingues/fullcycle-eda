package database

import (
	"database/sql"
	"testing"

	"github.com/lucasmdomingues/wallet-core/internal/domain/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type CustomerDBTestSuite struct {
	suite.Suite

	db         *sql.DB
	repository *CustomerDB
}

func (suite *CustomerDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	suite.db = db

	_, err = db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at DATE)")
	suite.NoError(err)

	suite.repository = NewCustomerDB(db)
}

func (suite *CustomerDBTestSuite) TearDownSuite() {
	suite.db.Exec("DROP TABLE customers")

	err := suite.db.Close()
	suite.NoError(err)
}

func (suite *CustomerDBTestSuite) TestSave() {
	customer, err := entity.NewCustomer("John Doe", "j@j")
	suite.NoError(err)

	err = suite.repository.Save(customer)
	suite.NoError(err)
}

func (suite *CustomerDBTestSuite) TestGet() {
	customer, err := entity.NewCustomer("John Doe", "j@j")
	suite.NoError(err)

	err = suite.repository.Save(customer)
	suite.NoError(err)

	got, err := suite.repository.Get(customer.ID)
	suite.NoError(err)

	suite.Equal(customer.ID, got.ID)
	suite.Equal(customer.Name, got.Name)
	suite.Equal(customer.Email, got.Email)
}

func TestCustomerDBTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerDBTestSuite))
}
