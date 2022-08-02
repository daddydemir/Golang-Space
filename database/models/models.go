package models

type Users struct {
	Id       int
	Name     string
	Surname  string
	Password string
}

type Customers struct {
	Userid int
	Phone  string
	People Users `gorm:"foreignKey:userid;references:id"`
}
