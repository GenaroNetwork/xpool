package models

import (
	"github.com/jinzhu/gorm"
	"xpool/database"
)

type Token struct {
	gorm.Model
	Email	string	`json:"email"`
	TokenRes	string	`gorm:"index:Token"`
	Timestamp int64     `json:"timestamp"`
	Role        int		`json:"role"`
}

func init() {
	db := database.GetDB()
	db.AutoMigrate(&Token{})
}


func SaveToken(token *Token)  {
	db := database.GetDB()
	db.Create(token)
}

func GetEmailByToken(token string) Token {
	var tokenRes Token
	db := database.GetDB()
	db.Where("token_res = ?",token).Last(&tokenRes)
	return tokenRes
}


func DeleteToken( email string)  {
	var token Token
	db := database.GetDB()
	db.Delete(&token,"email = ?", email)
}