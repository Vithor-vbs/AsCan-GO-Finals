package models

type User struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Username   string `json:"username" gorm:"unique"`
	FirstName  string `json:"firsName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"userStatus"`
}
