package config

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func ConnectToDB() {
	db, err := sql.Open(DbDriver, DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	log.Println("Connected to database")
}
