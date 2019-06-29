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
	Tag int
}

func init() {
	db := database.GetDB()
	db.AutoMigrate(&AutoTransaction{})
}

func SaveAutoTransaction (autoTransaction *AutoTransaction) error {
	db := database.GetDB()
	return db.Save(autoTransaction).Error
}


func GetAutoTransaction(status int) []AutoTransaction {
	var autoTransaction []AutoTransaction
	db := database.GetDB()
	db.Where("status = ? and tag = ?",status,0).Find(&autoTransaction)
	return autoTransaction
}

func UpdateAutoTransaction(hash string)  {
	db := database.GetDB()
	db.Model(&AutoTransaction{}).Where("hash = ?", hash).Updates(
		map[string]interface{}{"status": 1})
}

func UpdateAutoTransactionTag(status int)  {
	db := database.GetDB()
	db.Model(&AutoTransaction{}).Where("status = ? and tag = ?",status,0).Updates(
		map[string]interface{}{"tag": 1})
}