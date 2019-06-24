package models

import (
	"github.com/jinzhu/gorm"
	"xpool/database"
)

type Income struct {
	gorm.Model
	Email string
	TotalIncome float64
	IncomeBalance    float64
	UpdateUser uint
}

type IncomeLog struct {
	gorm.Model
	Email string
	TotalIncome float64
	IncomeBalance    float64
	UpdateUser uint
	Err string
}


func UpdateIncome(email string,totalIncome,incomeBalance float64,updateUser uint,operating string) bool {
	tx := database.GetDB()
	db := tx.Begin()
	var err error
	if "create" == operating  {
		err = db.Create(&Income{
			Email: email,
			TotalIncome: totalIncome,
			IncomeBalance:incomeBalance,
			UpdateUser:updateUser,
		}).Error
	}else if "update" == operating {
		err = db.Model(&Income{}).Where("email = ?", email).Updates(
			map[string]interface{}{"income_balance": incomeBalance,"total_income": totalIncome,"update_user":updateUser}).Error
	}

	err = db.Create(&IncomeLog{
		Email: email,
		TotalIncome: totalIncome,
		IncomeBalance:incomeBalance,
		UpdateUser:updateUser,
		Err:err.Error(),
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}
	db.Commit()
	return true
}


func GetIncomeInfoById(email string) Income {
	var income Income
	db := database.GetDB()
	db.Where("email = ?",email).Last(&income)
	return income
}