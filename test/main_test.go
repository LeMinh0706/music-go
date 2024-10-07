package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/LeMinh0706/music-go/db"
	"github.com/LeMinh0706/music-go/util"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("Can load config: ", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect db:", err)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
