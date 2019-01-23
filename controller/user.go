package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"xpool/models"
)

var User user = user{}
type user struct{}


func (u *user) CreateUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	verificationcode := c.PostForm("code")
	adderss := c.PostForm("adderss")
	c.JSON(http.StatusOK,CreateUserServices(email,password,verificationcode,adderss))
}


func (u *user) GetVerificationCode(c *gin.Context) {
	email := c.PostForm("email")
	c.JSON(http.StatusOK,GetVerificationCodeServices(email))
}

func (u *user) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	c.JSON(http.StatusOK,LoginServices(email,password))
}

func (u *user) GetUserByToken(c *gin.Context)  {
	token := c.PostForm("token")
	c.JSON(http.StatusOK,GetUserByTokenServices(token))
}

func (u *user) ForgetPassword (c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	verificationcode := c.PostForm("code")
	c.JSON(http.StatusOK,ForgetPasswordServices(email,password,verificationcode))
}

func (u *user)ResetPassword (c *gin.Context) {
	token := c.PostForm("token")
	password := c.PostForm("password")
	newPassword := c.PostForm("newPassword")
	c.JSON(http.StatusOK,ResetPasswordServices(token,password,newPassword))
}

func CreateUserServices(email,password,code,adderss string) Response  {
	emailVerify :=  VerifyEmailFormat(email)
	if true != emailVerify {
		return ResponseFun("email 格式错误",10000)
	}

	adderssVerify :=  VerifyEthAdderss(adderss)
	if true != adderssVerify {
		return ResponseFun("gnx address 格式错误",10001)
	}


	getUser :=  models.GetUserByEmail(email)
	if "" != getUser.Email {
		return ResponseFun("email 已存在",10002)
	}

	getUser =  models.GetUserByEthAdderss(adderss)
	if "" != getUser.Adderss {
		return ResponseFun("gnx address 已存在",10003)
	}


	verificationCode := models.GetVerificationCodeByEmail(email)
	if code == verificationCode.Code && time.Now().Unix() < verificationCode.Timestamp + 120 {
		DeleteVerificationCode(verificationCode.Code)
	}else {
		return ResponseFun("验证码错误或已过期",10004)
	}

	saltValue := GetRandomString(10)
	if 6 > len(password) {
		return ResponseFun("password 长度应大于5位",10006)
	}
	passwordVal := MD5(saltValue+MD5(password+saltValue)+saltValue)
	models.SaveUser(&models.User{
		Email:email,
		Password:passwordVal,
		SaltValue:saltValue,
	})
	return ResponseFun("注册成功",200)
}

func GetVerificationCodeServices(email string)	Response {
	emailVerify :=  VerifyEmailFormat(email)
	if true != emailVerify {
		return ResponseFun("email 格式错误",10008)
	}
	verificationCode := models.GetVerificationCodeByEmail(email)
	if time.Now().Unix() < verificationCode.Timestamp + 120 {
		result := MailTemplate(verificationCode.Code,email)
		if true == result {
			return ResponseFun("邮件发送成功",200)
		}
		return ResponseFun("邮件发送失败",10010)
	}
	code := GetRandomString(5)
	result := MailTemplate(code,email)
	if true == result {
		models.SaveVerificationCode(&models.VerificationCode{
			Code:code,
			Timestamp:time.Now().Unix(),
			Email:email,
		})
		return ResponseFun("邮件发送成功",200)
	}
	return ResponseFun("邮件发送失败",10012)
}

func LoginServices(email,password string) Response {
	getUser :=  models.GetUserByEmail(email)
	if "" == getUser.SaltValue {
		return ResponseFun("登录失败",10014)
	}

	passwordVal := MD5(getUser.SaltValue+MD5(password+getUser.SaltValue)+getUser.SaltValue)
	if passwordVal != getUser.Password {
		return ResponseFun("登录失败",10016)
	}
	models.DeleteToken(email)
	token := models.Token{
		Timestamp:time.Now().Unix(),
		Email:email,
		TokenRes:GetRandomString(50),
	}
	models.SaveToken(&token)
	return ResponseFun(token,200)
}

type UserInfo struct {
	Email	string `json:"email"`
	Address string	`json:"address"`
}

func GetUserByTokenServices(token string)  Response {
	if 50 != len(token) {
		return ResponseFun("获取用户信息失败",10018)
	}

	result := models.GetEmailByToken(token)
	if time.Now().Unix() < result.Timestamp + 43200 {
		user := models.GetUserByEmail(result.Email)
		return ResponseFun(UserInfo{Email:user.Email},200)
	}
	return ResponseFun("获取用户信息失败",10020)
}

func GetUserInfoByToken(token string) UserInfo {
	result := models.GetEmailByToken(token)
	if time.Now().Unix() < result.Timestamp + 43200 {
		user := models.GetUserByEmail(result.Email)
		return UserInfo{Email:user.Email,Address:user.Adderss}
	}
	return UserInfo{}
}

func ForgetPasswordServices(email,password,code string) Response  {
	emailVerify :=  VerifyEmailFormat(email)
	if true != emailVerify {
		return ResponseFun("email 格式错误",10022)
	}
	getUser :=  models.GetUserByEmail(email)
	if "" == getUser.Email {
		return ResponseFun("密码找回失败",10024)
	}

	verificationCode := models.GetVerificationCodeByEmail(email)
	if code == verificationCode.Code && time.Now().Unix() < verificationCode.Timestamp + 120 {
		DeleteVerificationCode(verificationCode.Code)
	}else {
		return ResponseFun("验证码错误或已过期",10026)
	}

	saltValue := GetRandomString(10)
	if 6 > len(password) {
		return ResponseFun("password 长度应大于5位",10028)
	}
	passwordVal := MD5(saltValue+MD5(password+saltValue)+saltValue)
	models.UpdateUser(email,saltValue,passwordVal)
	return ResponseFun("密码找回成功",200)
}

func DeleteVerificationCode(code string)  {
	models.DeleteVerificationCode(code)
}


func ResetPasswordServices(token,password,newPassword string) Response  {
	if 6 > len(newPassword) {
		return ResponseFun("newPassword 长度应大于5位",10030)
	}
	UserInfo :=  GetUserInfoByToken(token)
	if "" == UserInfo.Email {
		return ResponseFun("token 无效",10032)
	}
	getUser :=  models.GetUserByEmail(UserInfo.Email)
	passwordVal := MD5(getUser.SaltValue+MD5(password+getUser.SaltValue)+getUser.SaltValue)
	if passwordVal != getUser.Password {
		return ResponseFun("原始密码错误",10033)
	}
	saltValue := GetRandomString(10)
	newPasswordVal := MD5(saltValue+MD5(newPassword+saltValue)+saltValue)
	models.UpdateUser(UserInfo.Email,saltValue,newPasswordVal)
	return ResponseFun("重置成功",200)
}

func CheckPassword(token,password string) bool {
	UserInfo :=  GetUserInfoByToken(token)
	if "" == UserInfo.Email {
		return false
	}
	getUser :=  models.GetUserByEmail(UserInfo.Email)
	passwordVal := MD5(getUser.SaltValue+MD5(password+getUser.SaltValue)+getUser.SaltValue)
	if passwordVal != getUser.Password {
		return false
	}
	return true
}