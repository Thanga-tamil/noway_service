package config

import (
	"database/sql"
	"log"

	// Import the SQLite driver
	// The underscore import registers the driver with database/sql
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitSqlite() {

	log.Println("Initialize sqlite db")

	sqlite, err := sql.Open("sqlite3", "chats.db")
	DB = sqlite

	if err != nil {
		log.Println("Error while opening session with sqlite: ", err)
		panic(err)
	}

}
