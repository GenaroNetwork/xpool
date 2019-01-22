package models

import (
	"github.com/jinzhu/gorm"
	"xpool/database"
)

// State 1 待审核 3 审核通过   5 审核拒绝
type Deposit struct {
	gorm.Model
	State	int
	Email	string
	Hash    string	`gorm:"index:Hash"`
	Reason	string
	Value 	float64
}


func SaveDeposit(deposit *Deposit)  {
	db := database.GetDB()
	db.Create(deposit)
}


func GetDepositInfoByHsah(hsah string) Deposit {
	var deposit Deposit
	db := database.GetDB()
	db.Where("hsah = ?",hsah).Last(&deposit)
	return deposit
}