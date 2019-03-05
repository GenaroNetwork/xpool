package models

import (
	"github.com/jinzhu/gorm"
	"xpool/database"
)



// State 1 待审核 3 审核通过   5 审核拒绝
type LoanMining struct {
	gorm.Model
	State	int
	Email	string
	Loan    float64
	Reason	string
	Deposit 	float64
	UpdateUser uint
}


type LoanMiningLog struct {
	gorm.Model
	State int
	Deposit float64
	Reason,Email string
	UpdateUser	uint
	Balance   float64
	LogType		int
	Loan     float64
}


type UserLoanMiningBalance struct {
	gorm.Model
	Email string
	Deposit float64
	Loan    float64
	State int
	UpdateUser uint
	Address	string
}

type ExtractLoanMiningBalance struct {
	gorm.Model
	Email string
	Deposit float64
	Loan    float64
	State int
	UpdateUser uint
	Address	string
	Reason string
	DepositId uint
}



func GetUserLoanMiningBalanceByEmail(email string) UserLoanMiningBalance {
	var userLoanMiningBalance UserLoanMiningBalance
	db := database.GetDB()
	db.Where("email = ?",email).Last(&userLoanMiningBalance)
	return userLoanMiningBalance
}

func SaveLoanMining(state int, email string, value,balance float64,update_user uint,loan float64) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Create(&LoanMining{
		State:state,
		Email:email,
		Deposit:value,
		Loan:loan,
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

	err = db.Create(&LoanMiningLog{
		State:state,
		Deposit:value,
		Reason:"",
		Email:email,
		UpdateUser:update_user,
		Balance:balance,
		LogType:1,
		Loan:loan,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}


func GetLoanMiningInfoById(id string) LoanMining {
	var loanMining LoanMining
	db := database.GetDB()
	db.Where("id = ?",id).Last(&loanMining)
	return loanMining
}



func UpdateLoanMining(state int,deposit float64,reason,email string,depositId,update_user uint,loan float64,address,operating string) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Model(&LoanMining{}).Where("email = ? and id = ?", email,depositId).Updates(
		map[string]interface{}{"state": state, "reason": reason,"update_user":update_user}).Error

	if nil != err {
		db.Rollback()
		return false
	}


	if "create" == operating && 3 == state {
		err = db.Create(&UserLoanMiningBalance{
			Email: email,
			Deposit: deposit,
			Loan:loan,
			State:1,
			Address:address,
			UpdateUser: update_user,
		}).Error
	}else if "update" == operating && 3 == state {
		err = db.Model(&UserLoanMiningBalance{}).Where("email = ?", email).Updates(
			map[string]interface{}{"deposit": deposit,"loan": loan,"update_user":update_user}).Error
	}

	if nil != err {
		db.Rollback()
		return false
	}

	if 5 == state {
		err = db.Model(&UserDepositBalance{}).Where("email = ?", email).Updates(
			map[string]interface{}{"balance": deposit,"update_user":update_user}).Error
	}

	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&LoanMiningLog{
		State:state,
		Deposit:deposit,
		Reason:reason,
		Email:email,
		UpdateUser:update_user,
		LogType:0,
		Loan:loan,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}

func ExtractLoanMining(depositId uint,loan,deposit float64,email string,update_user uint) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Model(&UserLoanMiningBalance{}).Where("email = ? and id = ?", email,depositId).Updates(
		map[string]interface{}{"loan": 0, "deposit": 0,"update_user":update_user}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&ExtractLoanMiningBalance{
		Email: email,
		Deposit: deposit,
		Loan:loan,
		State:1,
		DepositId:depositId,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&LoanMiningLog{
		State:1,
		Deposit:deposit,
		Reason:"",
		Email:email,
		UpdateUser:update_user,
		LogType:1,
		Loan:loan,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}


func GetExtractLoanMiningBalanceInfoById(id string) ExtractLoanMiningBalance {
	var extractLoanMiningBalance ExtractLoanMiningBalance
	db := database.GetDB()
	db.Where("id = ?",id).Last(&extractLoanMiningBalance)
	return extractLoanMiningBalance
}


func UpdateExtractLoanMining(state int,deposit,loan float64,reason,email string,depositId,update_user uint) bool {
	tx := database.GetDB()
	db := tx.Begin()

	err := db.Model(&ExtractLoanMiningBalance{}).Where("email = ? and id = ?", email,depositId).Updates(
		map[string]interface{}{"state": state, "reason": reason,"update_user":update_user}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	if 3 == state {
		err = db.Model(&UserDepositBalance{}).Where("email = ?", email).Updates(
			map[string]interface{}{"balance": deposit,"update_user":update_user}).Error
	}else if 5 == state {
		err = db.Model(&UserLoanMiningBalance{}).Where("email = ? and id = ?", email,depositId).Updates(
			map[string]interface{}{"loan": loan, "deposit": deposit,"update_user":update_user}).Error
	}

	if nil != err {
		db.Rollback()
		return false
	}

	err = db.Create(&LoanMiningLog{
		State:state,
		Deposit:deposit,
		Reason:reason,
		Email:email,
		UpdateUser:update_user,
		LogType:2,
		Loan:loan,
	}).Error

	if nil != err {
		db.Rollback()
		return false
	}

	db.Commit()
	return true
}



func GetLoanMiningListByEmail(email string,page,pageSize int) []LoanMining {
	var loanMining []LoanMining
	db := database.GetDB()
	db.Where("email = ?",email).Limit(pageSize).Offset((page - 1) * pageSize).Find(&loanMining)
	return loanMining
}


func GetLoanMiningListCountByEmail(email string) int {
	var count int
	db := database.GetDB()
	db.Model(&LoanMining{}).Where("email = ?",email).Count(&count)
	return count
}


func GetLoanMiningList(page,pageSize int) []LoanMining {
	var loanMining []LoanMining
	db := database.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&loanMining)
	return loanMining
}


func GetLoanMiningListCount() int {
	var count int
	db := database.GetDB()
	db.Model(&LoanMining{}).Count(&count)
	return count
}


func GetExtractLoanMiningBalanceListByEmail(email string,page,pageSize int) []ExtractLoanMiningBalance {
	var extractLoanMiningBalance []ExtractLoanMiningBalance
	db := database.GetDB()
	db.Where("email = ?",email).Limit(pageSize).Offset((page - 1) * pageSize).Find(&extractLoanMiningBalance)
	return extractLoanMiningBalance
}


func GetExtractLoanMiningBalanceListCountByEmail(email string) int {
	var count int
	db := database.GetDB()
	db.Model(&ExtractLoanMiningBalance{}).Where("email = ?",email).Count(&count)
	return count
}

func GetExtractLoanMiningBalanceList(page,pageSize int) []ExtractLoanMiningBalance {
	var extractLoanMiningBalance []ExtractLoanMiningBalance
	db := database.GetDB()
	db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&extractLoanMiningBalance)
	return extractLoanMiningBalance
}


func GetExtractLoanMiningBalanceListCount() int {
	var count int
	db := database.GetDB()
	db.Model(&ExtractLoanMiningBalance{}).Count(&count)
	return count
}