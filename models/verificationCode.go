package models

import (
	"github.com/jinzhu/gorm"
	"xpool/database"
)

type VerificationCode struct {
	gorm.Model
	Email     string `gorm:"index:email"`
	Code      string `json:"code"`
	Timestamp int64  `json:"timestamp"`
}

func init() {
	db := database.GetDB()
	db.AutoMigrate(&VerificationCode{})
}

func SaveVerificationCode(verificationCode *VerificationCode) {
	db := database.GetDB()
	db.Create(verificationCode)
}

func GetVerificationCodeByEmail(email string) VerificationCode {
	var verificationCode VerificationCode
	db := database.GetDB()
	db.Where("email = ?", email).Last(&verificationCode)
	return verificationCode
}

func DeleteVerificationCode(code string) {
	var verificationCode VerificationCode
	db := database.GetDB()
	db.Delete(&verificationCode, "code = ?", code)
}
