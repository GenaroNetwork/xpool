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
	UpdateUser uint
}

type UserDepositBalance struct {
	gorm.Model
	Email string
	Balance float64
	UpdateUser uint
}

type DepositOperatingLog struct {
	gorm.Model
	State int
	Value float64
	Reason,Email string
	UpdateUser	uint
	Balance   float64
	LogType		int
}

// State 1 待审核 3 审核通过   5 审核拒绝
type ExtractDeposit struct {
	gorm.Model
	State	int
	Email	string
	Reason	string
	Value 	float64
	UpdateUser uint
}

func init() {
	db := database.GetDB()
	db.AutoMigrate(&Deposit{})
	db.AutoMigrate(&UserDepositBalance{})
	db.AutoMigrate(&DepositOperatingLog{})
	db.AutoMigrate(&ExtractDeposit{})
}

func SaveDeposit(deposit *Deposit)  {
	db := database.GetDB()
	db.Create(deposit)
}


func GetDepositInfoByHsah(hsah string) Deposit {
	var deposit Deposit
	db := database.GetDB()
	db.Where("hash = ?",hsah).Last(&deposit)
	return deposit
}


func GetDepositListByEmail(email string,page,pageSize int) []Deposit {
	var deposit []Deposit
	db := database.GetDB()
	db.Where("email = ?",email).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&deposit)
	return deposit
}

func GetDepositList(page,pageSize int) []Deposit {
	var deposit []Deposit
	db := database.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&deposit)
	return deposit
}

func GetDepositCountByEmail(email string) int {
	var count int
	db := database.GetDB()
	db.Model(&Deposit{}).Where("email = ?",email).Count(&count)
	return count
}

func GetDepositCount() int {
	var count int
	db := database.GetDB()
	db.Model(&Deposit{}).Count(&count)
	return count
}


func GetDepositInfoById(id string) Deposit {
	var deposit Deposit
	db := database.GetDB()
	db.Where("id = ?",id).Last(&deposit)
	return deposit
}


func GetUserDepositBalanceByEmail(email string) UserDepositBalance {
	var userDepositBalance UserDepositBalance
	db := database.GetDB()
	db.Where("email = ?",email).Last(&userDepositBalance)
	return userDepositBalance
}




func UpdateDeposit(state int,value float64,reason,email string,depositId,update_user uint,operating string) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Model(&Deposit{}).Where("email = ? and id = ?", email,depositId).Updates(
		map[string]interface{}{"state": state, "reason": reason,"update_user":update_user}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	if "create" == operating && 3 == state {
		err = db.Create(&UserDepositBalance{
			Email: email,
			Balance: value,
			UpdateUser: update_user,
		}).Error
	}else if "update" == operating && 3 == state {
		err = db.Model(&UserDepositBalance{}).Where("email = ?", email).Updates(
			map[string]interface{}{"balance": value,"update_user":update_user}).Error
	}

	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&DepositOperatingLog{
		State:state,
		Value:value,
		Reason:reason,
		Email:email,
		UpdateUser:update_user,
		LogType:0,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}

func SaveExtractDeposit(state int, email string, value,balance float64,update_user uint) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Create(&ExtractDeposit{
		State:state,
		Email:email,
		Value:value,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}
	err = db.Model(&UserDepositBalance{}).Where("email = ?", email).Updates(
		map[string]interface{}{"balance": balance,"update_user":update_user}).Error
	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&DepositOperatingLog{
		State:state,
		Value:value,
		Reason:"",
		Email:email,
		UpdateUser:update_user,
		Balance:balance,
		LogType:1,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}


func GetExtractDepositInfoById(id string) ExtractDeposit {
	var extractDeposit ExtractDeposit
	db := database.GetDB()
	db.Where("id = ?",id).Last(&extractDeposit)
	return extractDeposit
}


func UpdateExtractDeposit(state int,value float64,reason,email string,depositId,update_user uint) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Model(&ExtractDeposit{}).Where("email = ? and id = ?", email,depositId).Updates(
		map[string]interface{}{"state": state, "reason": reason,"update_user":update_user}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	if  5 == state {
		err = db.Model(&UserDepositBalance{}).Where("email = ?", email).Updates(
			map[string]interface{}{"balance": value,"update_user":update_user}).Error
	}

	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&DepositOperatingLog{
		State:state,
		Value:value,
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


func GetExtractDepositListByEmail(email string,page,pageSize int) []ExtractDeposit {
	var extractDeposit []ExtractDeposit
	db := database.GetDB()
	db.Where("email = ?",email).Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&extractDeposit)
	return extractDeposit
}


func GetExtractDepositCountByEmail(email string) int {
	var count int
	db := database.GetDB()
	db.Model(&ExtractDeposit{}).Where("email = ?",email).Count(&count)
	return count
}


func GetExtractDepositList(page,pageSize int) []ExtractDeposit {
	var extractDeposit []ExtractDeposit
	db := database.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&extractDeposit)
	return extractDeposit
}


func GetExtractDepositCount() int {
	var count int
	db := database.GetDB()
	db.Model(&ExtractDeposit{}).Count(&count)
	return count
}