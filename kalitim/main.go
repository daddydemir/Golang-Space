package main

import (
	"fmt"
	"kalitim/models"
)

func main() {
	var user models.User = models.User{Id: 1, Name: "Ali", Surname: "Test"}
	user.Name = "Mehmet"

	var ogretmen models.Teacher
	ogretmen.People = user

	fmt.Println(ogretmen.Branch)
}
