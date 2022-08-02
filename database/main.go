package main

import (
	"database/config"
	"database/models"
	"fmt"
	"time"

	"gorm.io/gorm/clause"
)

func main() {
	config.GetConnect()

	// var u Users
	// _ = config.DB.Find(&u)

	// fmt.Println(u)

	var f models.Customers
	_ = config.DB.Preload(clause.Associations).Find(&f)
	// Preload metodunu şu anda doğru kullanıyorum
	fmt.Println(f.User)
}

type Users struct {
	Id          int
	Name        string
	Surname     string
	Username    string
	Password    string
	Email       string
	CreatedDate time.Time
}
type Freelancers struct {
	UserId       int
	About        string
	ImagePath    string
	appellation  string
	averageScore float64
	User         Users
	// 41. satırdaki kodu ekleyince db'de bu şekilde bir Field yok şeklinde hata fırlatıyor.
}
