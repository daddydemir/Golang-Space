package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartDatabaseConnection() {

	dsn := "localhost user=postgres password=password123 dbname=databaseName port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error: ", err)
	}
	DB = db

	connection, _ := DB.DB()

	connection.Close()

}
