package test

import (
	"database/sql"
	"github.com/yemrealtanay/sms_templates/db/sqlc"
	"github.com/yemrealtanay/sms_templates/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
