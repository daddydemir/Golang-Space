package models

import "gorm.io/gorm"

type User struct {
	//gorm.Model
	Id       int `gorm:"primaryKey"`
	Username string
	Orders   Order
	//Status   Status
}

type Order struct {
	Id       int  `gorm:"primaryKey"`
	UserId   uint `gorm:"foreignKey"`
	StatusId int  `gorm:"foreignKey"`
	Price    float32
	Status   Status
}

type Status struct {
	StatusId int `gorm:"primaryKey"`
	Content  string
}
type Test struct {
	gorm.Model
	Code  string
	Price uint
}
