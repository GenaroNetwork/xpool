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
	Adderss		string	`json:"adderss"`
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

func UpdateUser(email,saltValue, password string)  {
	var user User
	db := database.GetDB()
	db.Model(&user).Where("email = ?", email).Updates(map[string]interface{}{"salt_value": saltValue, "password": password})
}