package main

import (
	"fmt"
	"nsted/config"
	"nsted/models"
)

func main() {
	config.GetConnect()
	//config.DB.AutoMigrate(&models.Status{})
	//config.DB.AutoMigrate(&models.Order{})
	//config.DB.AutoMigrate(&models.User{})
	var a []models.User
	fmt.Println("test")
	config.DB.Preload("Orders").Preload("Orders.Status").Find(&a)
	for _, model := range a {
		fmt.Println("NAME   : ", model.Username)
		fmt.Println("STATUS : ", model.Orders.Status.Content)
	}

	//var b models.Order
	//config.DB.Preload("Status").Find(&b)
	//fmt.Println("Data : ", b.Status.Content)
}
