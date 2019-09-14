package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xpool/models"
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


func (u *income) ExtractIncomeBalance(c *gin.Context) {
	value := c.PostForm("value")
	token := c.PostForm("token")
	password := c.PostForm("password")
	c.JSON(http.StatusOK,ExtractIncomeBalanceServices(token,value,password))
}


func (u *income) ExtractIncomeList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,ExtractIncomeListServices(page,pageSize,token))
}

func (u *income) AdminExtractIncomeList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,AdminExtractIncomeListServices(page,pageSize,token))
}


func (u *income) ExtractIncomeReview(c *gin.Context) {
	reviewId := c.PostForm("reviewId")
	reason := c.PostForm("reason")
	token := c.PostForm("token")
	password := c.PostForm("password")
	statesStr := c.PostForm("states")
	c.JSON(http.StatusOK,ExtractIncomeReviewServices(reviewId,reason,token,password,statesStr))
}

func IncomeTotalServices(token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",40000)
	}
	getIncomeInfoById := models.GetIncomeInfoById(userInfo.Email)
	return ResponseFun(Decimal(getIncomeInfoById.TotalIncome),200)
}

func IncomeBalanceServices(token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",40002)
	}
	getIncomeInfoById := models.GetIncomeInfoById(userInfo.Email)
	return ResponseFun(Decimal(getIncomeInfoById.IncomeBalance),200)
}



func ExtractIncomeBalanceServices(token,valueStr,password string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Address {
		return ResponseFun("获取地址失败",20026)
	}

	if !CheckPassword(token,password) {
		return ResponseFun("密码错误",20028)
	}

	value,err := strconv.ParseFloat(valueStr,64)

	if nil != err {
		return ResponseFun("提取金额错误",20030)
	}

	GetIncomeInfo := models.GetIncomeInfoById(userInfo.Email)

	balance := Round(GetIncomeInfo.IncomeBalance - value,3)

	if 0 > balance {
		return ResponseFun("余额不足",20032)
	}

	result := models.SaveIncome(1, userInfo.Email, value,balance,userInfo.Id)
	if true != result {
		return ResponseFun("申请提取余额失败",20034)
	}
	return ResponseFun("申请提取余额成功",200)
}



type IncomeList struct {
	ExtractIncome []models.ExtractIncome   `json:"extract_income"`
	Page	int	`json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

func ExtractIncomeListServices(pageStr,pageSizeStr,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",20014)
	}
	page,err:=strconv.Atoi(pageStr)
	if nil != err {
		page = 1
	}
	pageSize,err:=strconv.Atoi(pageSizeStr)
	if nil != err {
		pageSize = 100
	}

	if 0 >= page {
		page = 1
	}

	if 100 < pageSize {
		pageSize = 100
	}

	return ResponseFun(IncomeList{
		ExtractIncome:models.GetExtractIncomeListByEmail(userInfo.Email,page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.GetExtractIncomeCountByEmail(userInfo.Email),
	},200)
}


func AdminExtractIncomeListServices(pageStr,pageSizeStr,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",20014)
	}
	page,err:=strconv.Atoi(pageStr)
	if nil != err {
		page = 1
	}
	pageSize,err:=strconv.Atoi(pageSizeStr)
	if nil != err {
		pageSize = 100
	}

	if 0 >= page {
		page = 1
	}

	if 100 < pageSize {
		pageSize = 100
	}

	return ResponseFun(IncomeList{
		ExtractIncome:models.AdminGetExtractIncomeListByEmail(page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.AdminGetExtractIncomeCountByEmail(),
	},200)
}



func ExtractIncomeReviewServices(reviewId,reason,token,password,statesStr string) Response {
	userInfo := GetUserInfoByToken(token)
	states,err:=strconv.Atoi(statesStr)
	if nil != err {
		return ResponseFun("参数错误",20024)
	}

	if 3 != states && 5 != states {
		return ResponseFun("参数错误",20026)
	}
	if "" == userInfo.Email {
		return ResponseFun("token 无效",20016)
	}
	if !VerifyAdminRole(userInfo) {
		return ResponseFun("无权限操作",20018)
	}
	if !CheckPassword(token,password) {
		return ResponseFun("密码错误",20020)
	}
	extractIncomeInfo := models.GetExtractIncomeInfoById(reviewId)
	if 1 != extractIncomeInfo.State {
		return ResponseFun("操作错误",20022)
	}

	incomeInfo := models.GetIncomeInfoById(extractIncomeInfo.Email)
	var value float64
	if 5 == states {
		value = extractIncomeInfo.Value + incomeInfo.IncomeBalance
	}
	result := models.UpdateExtractIncome(states,value,reason,extractIncomeInfo.Email,extractIncomeInfo.ID,userInfo.Id)
	if true == result {
		return ResponseFun("审核成功",200)
	}else {
		return ResponseFun("审核失败",20024)
	}
}
