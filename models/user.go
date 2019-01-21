package models

import (
	"github.com/jinzhu/gorm"
	"xpool/database"
)

type User struct {
	gorm.Model
	Email	string	`gorm:"index:email"`
	SaltValue	string	`json:"salt_value"`
	Password	string	`json:"password"`
}

func SaveUser(user *User)  {
	db := database.GetDB()
	db.Create(user)
}


func GetUserByEmail(email string) User {
	var user User
	db := database.GetDB()
	db.Where("email = ?",email).Last(&user)
	return user
}