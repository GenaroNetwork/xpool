package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Income income = income{}
type income struct{}

func (u *income) IncomeTotal(c *gin.Context) {
	token := c.PostForm("token")
	c.JSON(http.StatusOK,IncomeTotalServices(token))
}

func (u *income) IncomeBalance(c *gin.Context) {
	token := c.PostForm("token")
	c.JSON(http.StatusOK,IncomeBalanceServices(token))
}



func IncomeTotalServices(token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",40000)
	}

	return ResponseFun(8000,200)
}

func IncomeBalanceServices(token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",40002)
	}
	return ResponseFun(3340,200)
}