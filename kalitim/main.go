package main

import (
	"fmt"
	"kalitim/models"
)

var sayi int = 10

func main() {
	var user models.User = models.User{Id: 1, Name: "Ali", Surname: "Test"}
	user.Name = "Mehmet"

	var ogretmen models.Teacher
	ogretmen.People = user

	fmt.Println(ogretmen.Branch)
	println("selam")
}
