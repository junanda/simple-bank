package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	. "github.com/junanda/simple-bank/repository"
	ifc "github.com/junanda/simple-bank/repository/interface_repo"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:triadpass@localhost:5432/simple_bank?sslmode=disable"
)

var (
	testAccounts ifc.AccountRepository
	testEntry    ifc.EntryRepository
	testTranfer  ifc.TransferRepository
	testDB       *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	testAccounts = NewAccountRepository(testDB)
	testEntry = NewEntriesRepository(testDB)
	testTranfer = NewTransferImpl(testDB)

	os.Exit(m.Run())
}
