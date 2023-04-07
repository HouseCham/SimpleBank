package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/HouseCham/SimpleBank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")	// parent folder
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to the database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
