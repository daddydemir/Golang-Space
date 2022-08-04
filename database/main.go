package main

import (
	"database/config"
	"database/handler"
	"net/http"
)

func main() {
	config.GetConnect()
	server := &http.Server{
		Addr:    ":1337",
		Handler: handler.MainRouting(),
	}
	server.ListenAndServe()

	// var customer models.Customers

	// customer.Username = "daddydemir"
	// customer.User.Mail = "mail.com"
	// customer.User.Phone = "123456"
	// customer.User.Password = "my_pass"
	// customer.User.RoleId = 1
	// customer.User.ImagePath = "images"

	// result := config.DB.Create(&customer)

	// if result.Error != nil {
	// 	fmt.Println("Hata")
	// }

	//AddCustomer()

	// var u Users
	// _ = config.DB.Find(&u)

	// fmt.Println(u)
	/*
		var f models.Customers
		_ = config.DB.Preload(clause.Associations).Find(&f)
		fmt.Println(f.People)
	*/
}

//func AddCustomer() {
//	var c models.Customers
//	c.Phone = "1234567890"
//	c.People.Password = "mypass"
//	c.People.Name = "Kemal"
//	c.People.Surname = "Bey"
//	result := config.DB.Create(&c)
//	if result.Error == nil {
//		fmt.Println("Kod doğru çalıştı.")
//	} else {
//		fmt.Println(result.Error, " Ulan hata aldık iyi mi?")
//	}
//}

type User struct {
	Id        int
	RoleId    int
	ImagePath string
	Phone     string
	Password  string
	Mail      string
}

type Stationery struct {
	UserId      int
	CompanyName string
	AddressId   int
	Score       float32
}
