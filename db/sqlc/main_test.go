package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const(
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:qwerty@localhost:5433?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil{
		log.Fatal(err.Error())
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}