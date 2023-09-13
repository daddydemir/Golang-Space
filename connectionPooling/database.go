package main

import (
	"database/sql"
	"log"
	"time"
)

func connectWithStandartLibrary() *sql.DB {

	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetConnMaxLifetime(time.Hour)

	return db
}
