package repos

import (
	"database/config"
	"database/models"

	"gorm.io/gorm/clause"
)

// get users
func Init() {
	var user models.Users
	_ = config.DB.Find(&user, "id = ?", 1)
	println("UserName : ", user.Password)
}

func GetCustomer() {
	var c models.Customers
	_ = config.DB.Find(&c, "userid = ?", 1)

	println("C : ", c.Username)

}

func GetAll() models.Customers {
	var c models.Customers
	_ = config.DB.Preload(clause.Associations).Find(&c)
	return c
}
