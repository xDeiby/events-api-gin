package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDb() {
	var err error
	
	DB, err = sql.Open("sqlite", "events.db")

	if err != nil {
		fmt.Println(err)
		panic("could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTableSql := `
		CREATE TABLE IF NOT EXISTS events (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name VARCHAR(25) NOT NULL,
				description TEXT NOT NULL,
				location VARCHAR(30) NOT NULL,
				date_time DATETIME NOT NULL,
				user_id INTEGER
		)
	`

	_, err := DB.Exec(createEventsTableSql)

	if err != nil {
		fmt.Println(err)
		panic("could not create events table.")
	}
}