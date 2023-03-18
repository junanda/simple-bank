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
)

var testAccounts ifc.AccountRepository

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	testAccounts = NewAccountRepository(conn)

	os.Exit(m.Run())
}
