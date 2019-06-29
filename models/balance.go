package models

import (
	"github.com/jinzhu/gorm"
	"xpool/database"
)

type UserBalance struct {
	gorm.Model
	Email string
	Balance float64
	State	int
	UpdateUser uint
}

type  ExtractBalance struct {
	gorm.Model
	Email string
	Balance float64
	State	int
	Reason	string
	UpdateUser uint
}


type  ExtractBalanceLog struct {
	gorm.Model
	Email string
	Balance float64
	State	int
	Reason	string
	UpdateUser uint
	LogType int
}

func init() {
	db := database.GetDB()
	db.AutoMigrate(&UserBalance{})
	db.AutoMigrate(&ExtractBalance{})
	db.AutoMigrate(&ExtractBalanceLog{})
}

func GetUserBalanceByEmail(email string) UserBalance {
	var userBalance UserBalance
	db := database.GetDB()
	db.Where("email = ?",email).Last(&userBalance)
	return userBalance
}

func ExtractUserBalance(balance float64,email string, update_user uint) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Model(&UserBalance{}).Where("email = ?", email).Updates(
		map[string]interface{}{"balance": 0,"state":1,"update_user":update_user}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&ExtractBalance{
		Balance:balance,
		State:1,
		Email:email,
	}).Error

	if err != err {
		db.Rollback()
		return false
	}

	err = db.Create(&ExtractBalanceLog{
		State:1,
		Balance:balance,
		Email:email,
		UpdateUser:update_user,
		LogType:2,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}


func GetExtractUserBalanceInfoById(id string) ExtractBalance {
	var extractBalance ExtractBalance
	db := database.GetDB()
	db.Where("id = ?",id).Last(&extractBalance)
	return extractBalance
}



func UpdateExtractUserBalance(state int,balance float64,reason,email string,depositId,update_user uint) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Model(&ExtractBalance{}).Where("email = ? and id = ?", email,depositId).Updates(
		map[string]interface{}{"state": state, "reason": reason,"update_user":update_user}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Model(&UserBalance{}).Where("email = ?", email).Updates(
		map[string]interface{}{"state":0}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	if 5 == state {
		err = db.Model(&UserBalance{}).Where("email = ?", email).Updates(
			map[string]interface{}{"balance": balance,"update_user":update_user}).Error
	}

	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&ExtractBalanceLog{
		State:state,
		Balance:balance,
		Reason:reason,
		Email:email,
		UpdateUser:update_user,
		LogType:2,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}

func AddUserBalance(balance float64,email string,operating string) bool {
	db := database.GetDB()
	if "create" == operating {
		db.Create(&UserBalance{
			Email: email,
			Balance: balance,
			State: 0,
		})
	}else if "update" == operating {
		db.Model(&UserBalance{}).Where("email = ?", email).Updates(
			map[string]interface{}{"balance": balance})
	}

	db.Create(&ExtractBalanceLog{
		Balance:balance,
		Reason:operating,
		Email:email,
		LogType:0,
	})

	return true
}


func GetExtractBalanceListByEmail(email string,page,pageSize int) []ExtractBalance {
	var extractBalance []ExtractBalance
	db := database.GetDB()
	db.Where("email = ?",email).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&extractBalance)
	return extractBalance
}


func GetExtractBalanceListCountByEmail(email string) int {
	var count int
	db := database.GetDB()
	db.Model(&ExtractBalance{}).Where("email = ?",email).Count(&count)
	return count
}


func GetExtractBalanceList(page,pageSize int) []ExtractBalance {
	var extractBalance []ExtractBalance
	db := database.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&extractBalance)
	return extractBalance
}


func GetExtractBalanceListCount() int {
	var count int
	db := database.GetDB()
	db.Model(&ExtractBalance{}).Count(&count)
	return count
}