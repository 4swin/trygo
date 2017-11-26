// main_test.go

package main_test

import (
	"os"
	"testing"

	"."
)

var a main.App

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS employee`

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize("root", "", "127.0.0.1:3306", "test")
	//   a.Initialize(
	// os.Getenv("TEST_DB_USERNAME"),
	// os.Getenv("TEST_DB_PASSWORD"),
	// os.Getenv("TEST_DB_NAME"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func (a *App) Initialize(user, password, host, dbname string) {
	// connectionString := fmt.Sprintf("%s:%s@%s/%s", user, password, host, dbname)

	var err error
	// a.DB, err = sql.Open("mysql", connectionString)
	a.DB, err = sql.Open("mysql", "root:@/test")
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}
