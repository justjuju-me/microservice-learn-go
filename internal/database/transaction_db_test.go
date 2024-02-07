package database

import (
	"database/sql"
	"testing"

	"github.com/justjuju-me/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db *sql.DB
	client1 *entity.Client
	client2 *entity.Client
	accountFrom *entity.Account
	accountTo *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_from varchar(255), account_to varchar(255), amount int, created_at date)")
	client1, err := entity.NewClient("John Doe", "j@j.com")
	s.Nil(err)
	client2, err := entity.NewClient("Jane Doe", "x@x.com")
	s.Nil(err)

	// creating accounts
	accountFrom := entity.NewAccount(client1)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom
	accountTo := entity.NewAccount(client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
