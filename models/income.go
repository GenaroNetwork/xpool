package models

import (
	"github.com/jinzhu/gorm"
	"xpool/database"
)

type Income struct {
	gorm.Model
	Email         string
	TotalIncome   float64
	IncomeBalance float64
	UpdateUser    uint
}

type IncomeLog struct {
	gorm.Model
	Email         string
	TotalIncome   float64
	IncomeBalance float64
	UpdateUser    uint
}

// State 1 待审核 3 审核通过   5 审核拒绝
type ExtractIncome struct {
	gorm.Model
	State      int
	Email      string
	Reason     string
	Value      float64
	UpdateUser uint
}

type ExtractIncomeLog struct {
	gorm.Model
	State         int
	Value         float64
	Reason, Email string
	UpdateUser    uint
	Balance       float64
	LogType       int
}

func init() {
	db := database.GetDB()
	db.AutoMigrate(&Income{})
	db.AutoMigrate(&IncomeLog{})
	db.AutoMigrate(&ExtractIncome{})
	db.AutoMigrate(&ExtractIncomeLog{})
}

func UpdateIncome(email string, totalIncome, incomeBalance float64, updateUser uint, operating string) bool {
	tx := database.GetDB()
	db := tx.Begin()
	var err error
	if "create" == operating {
		err = db.Create(&Income{
			Email:         email,
			TotalIncome:   totalIncome,
			IncomeBalance: incomeBalance,
			UpdateUser:    updateUser,
		}).Error
	} else if "update" == operating {
		err = db.Model(&Income{}).Where("email = ?", email).Updates(
			map[string]interface{}{"income_balance": incomeBalance, "total_income": totalIncome, "update_user": updateUser}).Error
	}

	err = db.Create(&IncomeLog{
		Email:         email,
		TotalIncome:   totalIncome,
		IncomeBalance: incomeBalance,
		UpdateUser:    updateUser,
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
	db.Where("email = ?", email).Last(&income)
	return income
}

func SaveIncome(state int, email string, value, balance float64, update_user uint) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Create(&ExtractIncome{
		State: state,
		Email: email,
		Value: value,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}
	err = db.Model(&Income{}).Where("email = ?", email).Updates(
		map[string]interface{}{"income_balance": balance, "update_user": update_user}).Error
	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&ExtractIncomeLog{
		State:      state,
		Value:      value,
		Reason:     "",
		Email:      email,
		UpdateUser: update_user,
		Balance:    balance,
		LogType:    1,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}

func GetExtractIncomeListByEmail(email string, page, pageSize int) []ExtractIncome {
	var extractIncome []ExtractIncome
	db := database.GetDB()
	db.Where("email = ?", email).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&extractIncome)
	return extractIncome
}

func GetExtractIncomeCountByEmail(email string) int {
	var count int
	db := database.GetDB()
	db.Model(&ExtractIncome{}).Where("email = ?", email).Count(&count)
	return count
}

func AdminGetExtractIncomeListByEmail(page, pageSize int) []ExtractIncome {
	var extractIncome []ExtractIncome
	db := database.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&extractIncome)
	return extractIncome
}

func AdminGetExtractIncomeCountByEmail() int {
	var count int
	db := database.GetDB()
	db.Model(&ExtractIncome{}).Count(&count)
	return count
}

func GetExtractIncomeInfoById(id string) ExtractIncome {
	var extractIncome ExtractIncome
	db := database.GetDB()
	db.Where("id = ?", id).Last(&extractIncome)
	return extractIncome
}

func UpdateExtractIncome(state int, value float64, reason, email string, depositId, update_user uint) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Model(&ExtractIncome{}).Where("email = ? and id = ?", email, depositId).Updates(
		map[string]interface{}{"state": state, "reason": reason, "update_user": update_user}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	if 5 == state {
		err = db.Model(&Income{}).Where("email = ?", email).Updates(
			map[string]interface{}{"income_balance": value, "update_user": update_user}).Error
	}

	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&ExtractIncomeLog{
		State:      state,
		Value:      value,
		Reason:     reason,
		Email:      email,
		UpdateUser: update_user,
		LogType:    2,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}
