package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error

	DB, err = sql.Open("sqlite3", "car.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {

	createUsersTable := `CREATE TABLE IF NOT EXISTS USERS(
	id INTEGER AUTOINCREEMENT, name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE, password TEXT NOT NULL,
	createdAt TIMESTAMP
	)  
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic(err)
	}
}
