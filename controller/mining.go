package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xpool/models"
	"strconv"
)

var Mining mining = mining{}
type mining struct{}

const LEVER  =  100000

func (u *mining) LoanMining(c *gin.Context) {
	value := c.PostForm("value")
	password := c.PostForm("password")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,LoanMiningServices(token,value,password))
}

func (u *mining) LoanMiningReview(c *gin.Context) {
	loanMiningId := c.PostForm("loanMiningId")
	reason := c.PostForm("reason")
	token := c.PostForm("token")
	password := c.PostForm("password")
	states := c.PostForm("states")

	c.JSON(http.StatusOK,LoanMiningReviewServices(loanMiningId,reason,token,password,states))
}


func LoanMiningServices(token,valueStr,password string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Address {
		return ResponseFun("获取地址失败",30000)
	}

	if !CheckPassword(token,password) {
		return ResponseFun("密码错误",30002)
	}

	value,err := strconv.ParseFloat(valueStr,64)

	if nil != err {
		return ResponseFun("申请挖币金额错误",30004)
	}

	userDepositBalanceInfo := models.GetUserDepositBalanceByEmail(userInfo.Email)

	balance := Round(userDepositBalanceInfo.Balance - value,3)

	if 0 > balance {
		return ResponseFun("保证金余额不足",30006)
	}

	userLoanMiningBalanceInfo := models.GetUserLoanMiningBalanceByEmail(userInfo.Email)
	if "" == userLoanMiningBalanceInfo.Email {
		if 500000 > value*LEVER {
			return ResponseFun("借贷金额不足50万，无法挖矿",30008)
		}
	}

	result := models.SaveLoanMining(1, userInfo.Email, value,balance,userInfo.Id,value*LEVER)
	if true != result {
		return ResponseFun("申请借币挖矿失败",30010)
	}

	return ResponseFun("申请借币挖矿成功",200)
}


func LoanMiningReviewServices(loanMiningId,reason,token,password,statesStr string) Response {
	userInfo := GetUserInfoByToken(token)
	states,err:=strconv.Atoi(statesStr)
	if nil != err {
		return ResponseFun("参数错误",30024)
	}

	if 3 != states && 5 != states {
		return ResponseFun("参数错误",30026)
	}
	if "" == userInfo.Email {
		return ResponseFun("token 无效",30016)
	}
	if !VerifyAdminRole(userInfo) {
		return ResponseFun("无权限操作",30018)
	}
	if !CheckPassword(token,password) {
		return ResponseFun("密码错误",30020)
	}
	depositInfo := models.GetLoanMiningInfoById(loanMiningId)
	if 1 != depositInfo.State {
		return ResponseFun("操作错误",30022)
	}

	userDepositBalanceInfo := models.GetUserLoanMiningBalanceByEmail(depositInfo.Email)
	var deposit,loan float64
	if 3 == states {
		deposit = depositInfo.Deposit + userDepositBalanceInfo.Deposit
		loan = depositInfo.Loan + userDepositBalanceInfo.Loan
	}
	if 5 == states {
		userDepositBalanceInfo := models.GetUserDepositBalanceByEmail(depositInfo.Email)
		deposit = depositInfo.Deposit + userDepositBalanceInfo.Balance
	}

	var result bool
	if "" == userDepositBalanceInfo.Email {
		result = models.UpdateLoanMining(states,deposit,reason,depositInfo.Email,depositInfo.ID,userInfo.Id,loan,"create")
	}else {
		result = models.UpdateLoanMining(states,deposit,reason,depositInfo.Email,depositInfo.ID,userInfo.Id,loan,"update")
	}
	if true == result {
		return ResponseFun("审核成功",200)
	}else {
		return ResponseFun("审核失败",30024)
	}
}