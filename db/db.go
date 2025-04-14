package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	fmt.Println("***A")
	DB, err := sql.Open("sqlite3", "api.db")
	fmt.Println("***B")
	if err != nil {
		fmt.Println("***C.error")
		panic("Could not connect to database.")
	}
	fmt.Println("***D")
	DB.SetMaxOpenConns(10)
	fmt.Println("***E")
	DB.SetMaxIdleConns(5)
	fmt.Println("***F")

	createTables()
	fmt.Println("***G")
}

func createTables() {
	fmt.Println("***CreateTables - A")
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	fmt.Println("***CreateTables - B")
	_, err := DB.Exec(createUsersTable)
	fmt.Println("***CreateTables - C")
	if err != nil {
		panic("Could not create users table.")
	}

	fmt.Println("***CreateTables - D")
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	fmt.Println("***CreateTables - E")
	_, err = DB.Exec(createEventsTable)
	fmt.Println("***CreateTables - F")
	if err != nil {
		fmt.Println("***CreateTables - ERROR")
		panic("Could not create events table.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}
}
