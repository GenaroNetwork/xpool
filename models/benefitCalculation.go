package models

import (
	"github.com/jinzhu/gorm"
	"xpool/database"
)

type AutoTransaction struct {
	gorm.Model
	Email string
	From string
	To 	string
	Value float64
	Hash string
	Status int
	Use int
}

func init() {
	db := database.GetDB()
	db.AutoMigrate(&AutoTransaction{})
}

func SaveAutoTransaction (autoTransaction *AutoTransaction) error {
	db := database.GetDB()
	return db.Save(autoTransaction).Error
}

