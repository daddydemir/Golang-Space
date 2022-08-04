package models

type Users struct {
	Id        int
	RoleId    int
	ImagePath string
	Phone     string
	Password  string
	Mail      string
}

//type Customers struct {
//	Userid int
//	Phone  string
//	People Users `gorm:"foreignKey:userid;references:id"`
//}

type Customers struct {
	User     Users  `json:"-" gorm:"foreignKey:user_id;references:id"`
	UserId   int    `json:"id"`
	Username string `json:"name"`
}

type Customer struct {
	UserId   int
	Username string
}
