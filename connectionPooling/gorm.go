package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func connectWithGorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open("username:password@tcp(localhost:3306)/test_db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(5)
	database.SetConnMaxIdleTime(time.Minute * 3)
	database.SetConnMaxLifetime(time.Hour * 1)

	return db
}
