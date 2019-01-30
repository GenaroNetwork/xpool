package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xpool/models"
	"strconv"
)



var Balance balance = balance{}
type balance struct{}

func (u *balance) ExtractBalance(c *gin.Context) {
	token := c.PostForm("token")
	password := c.PostForm("password")
	c.JSON(http.StatusOK,ExtractBalanceServices(token,password))
}

func (u *balance) ExtractBalanceReview(c *gin.Context)  {
	extractLoanMiningBalanceId := c.PostForm("reviewId")
	reason := c.PostForm("reason")
	token := c.PostForm("token")
	password := c.PostForm("password")
	statesStr := c.PostForm("states")
	c.JSON(http.StatusOK,ExtractBalanceReviewServices(extractLoanMiningBalanceId,reason,token,password,statesStr))
}

func (u *balance) ExtractBalanceList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,GetExtractBalanceListServices(page,pageSize,token))
}

func ExtractBalanceServices(token,password string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",40000)
	}

	if !CheckPassword(token,password) {
		return ResponseFun("密码错误",40002)
	}

	userBalanceInfo := models.GetUserBalanceByEmail(userInfo.Email)
	if 0 == userBalanceInfo.Balance  {
		return ResponseFun("余额为0",40004)
	}

	result := models.ExtractUserBalance(userBalanceInfo.Balance,userInfo.Email,userInfo.Id)
	if true != result {
		return ResponseFun("申请失败",40006)
	}
	return ResponseFun("申请成功",200)
}


func ExtractBalanceReviewServices(extractUserBalanceId,reason,token,password,statesStr string) Response {
	userInfo := GetUserInfoByToken(token)
	states,err:=strconv.Atoi(statesStr)
	if nil != err {
		return ResponseFun("参数错误",40008)
	}

	if 3 != states && 5 != states {
		return ResponseFun("参数错误",40010)
	}
	if "" == userInfo.Email {
		return ResponseFun("token 无效",40012)
	}
	if !VerifyAdminRole(userInfo) {
		return ResponseFun("无权限操作",40014)
	}
	if !CheckPassword(token,password) {
		return ResponseFun("密码错误",40016)
	}

	extractUserBalanceInfo := models.GetExtractUserBalanceInfoById(extractUserBalanceId)
	if 1 != extractUserBalanceInfo.State {
		return ResponseFun("操作错误",40018)
	}

	result := models.UpdateExtractUserBalance(states,extractUserBalanceInfo.Balance,reason,extractUserBalanceInfo.Email,extractUserBalanceInfo.ID,userInfo.Id)
	if true == result {
		return ResponseFun("审核成功",200)
	}else {
		return ResponseFun("审核失败",40020)
	}
}


func AddUserBalanceServices(balance float64,email string) bool {
	userBalanceInfo := models.GetUserBalanceByEmail(email)
	if 1 == userBalanceInfo.State  {
		return  false
	}

	if 0 == balance {
		return false
	}

	if "" == userBalanceInfo.Email {
		models.AddUserBalance(balance,email,"create")
	}else {
		models.AddUserBalance(balance,email,"update")
	}

	return true
}


type ExtractBalanceList struct {
	ExtractBalanceList []models.ExtractBalance   `json:"extract_balance_list"`
	Page	int	`json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

func GetExtractBalanceListServices(pageStr,pageSizeStr,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",40020)
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

	return ResponseFun(ExtractBalanceList{
		ExtractBalanceList:models.GetExtractBalanceListByEmail(userInfo.Email,page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.GetExtractBalanceListCountByEmail(userInfo.Email),
	},200)
}