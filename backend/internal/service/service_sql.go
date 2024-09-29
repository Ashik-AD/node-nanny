package service

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	CONN_STR = "file:../../database/store.db?cache=shared"
)

var db *sql.DB

func InitConnection() *sql.DB {
	if db == nil {
		conn, err := sql.Open("sqlite3", CONN_STR)
		if err != nil {
			log.Fatal("Failed to established connect: ", err)
		}

		pingErr := conn.Ping()
		if pingErr != nil {
			log.Fatal("Unabled to connect database: ", pingErr)
		}

		conn.SetMaxIdleConns(1)
		db = conn
	}
	return db
}
