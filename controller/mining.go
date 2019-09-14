package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xpool/models"
	"strconv"
)

var Mining mining = mining{}
type mining struct{}

func (u *mining) LoanMining(c *gin.Context) {
	value := c.PostForm("value")
	password := c.PostForm("password")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,LoanMiningServices(token,value,password))
}

func (u *mining) GetLoanMiningList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,GetLoanMiningListServices(page,pageSize,token))
}

func (u *mining) AdminGetLoanMiningList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,AdminGetLoanMiningListServices(page,pageSize,token))
}


func (u *mining) LoanMiningReview(c *gin.Context) {
	loanMiningId := c.PostForm("loanMiningId")
	reason := c.PostForm("reason")
	token := c.PostForm("token")
	password := c.PostForm("password")
	states := c.PostForm("states")
	address := c.PostForm("address")
	key := c.PostForm("key")
	pass := c.PostForm("pass")
	c.JSON(http.StatusOK,LoanMiningReviewServices(loanMiningId,reason,token,password,states,address,key, pass))
}

func (u *mining) IsBindingMiningAddress(c *gin.Context) {
	loanMiningId := c.PostForm("loanMiningId")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,IsBindingMiningAddressServices(loanMiningId,token))
}

func (u *mining) ExtractLoanMining(c *gin.Context) {
	token := c.PostForm("token")
	password := c.PostForm("password")
	c.JSON(http.StatusOK,ExtractLoanMiningServices(token,password))
}

func (u *mining) ExtractLoanMiningReview(c *gin.Context) {
	extractLoanMiningBalanceId := c.PostForm("reviewId")
	reason := c.PostForm("reason")
	token := c.PostForm("token")
	password := c.PostForm("password")
	statesStr := c.PostForm("states")
	c.JSON(http.StatusOK,ExtractLoanMiningReviewServices(extractLoanMiningBalanceId,reason,token,password,statesStr))
}

func (u *mining) GetExtractLoanMiningList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,ExtractLoanMiningListServices(page,pageSize,token))
}

func (u *mining) AdminGetExtractLoanMiningList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,AdminExtractLoanMiningListServices(page,pageSize,token))
}

func (u *mining) UserLoanMiningBalance(c *gin.Context) {
	token := c.PostForm("token")
	c.JSON(http.StatusOK,UserLoanMiningBalanceServices(token))
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
	//if "" == userLoanMiningBalanceInfo.Email {
	//	if MINING > value*LEVER {
	//		return ResponseFun("借贷金额不足"+strconv.Itoa(MINING)+"，无法挖矿",30008)
	//	}
	//}
	//
	//if MINING > value*LEVER + userLoanMiningBalanceInfo.Loan {
	//	return ResponseFun("借贷金额不足"+strconv.Itoa(MINING)+"，无法挖矿",30009)
	//}
	//
	//result := models.SaveLoanMining(1, userInfo.Email, value,balance,userInfo.Id,value*LEVER)

	if "" == userLoanMiningBalanceInfo.Email {
		if MINING > value {
			return ResponseFun("保证金不足"+strconv.Itoa(MINING)+"，无法挖矿",30008)
		}
	}

	if MINING > value + userLoanMiningBalanceInfo.Loan {
		return ResponseFun("保证金不足"+strconv.Itoa(MINING)+"，无法挖矿",30009)
	}

	result := models.SaveLoanMining(1, userInfo.Email, value,balance,userInfo.Id,value*LEVER)
	if true != result {
		return ResponseFun("申请借币挖矿失败",30010)
	}

	return ResponseFun("申请借币挖矿成功",200)
}


func LoanMiningReviewServices(loanMiningId,reason,token,password,statesStr,address,key, pass string) Response {
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
	if "" == userDepositBalanceInfo.Email {
		if true != VerifyEthAdderss(address) {
			return ResponseFun("gnx 挖矿地址格式错误",30026)
		}

		if "" == key {
			return ResponseFun("私钥错误",30027)
		}
		if "" == pass {
			return ResponseFun("挖矿密码错误",30030)
		}
	}

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
		result = models.UpdateLoanMining(states,deposit,reason,depositInfo.Email,depositInfo.ID,userInfo.Id,loan,address,"create",key, pass)
	}else {
		result = models.UpdateLoanMining(states,deposit,reason,depositInfo.Email,depositInfo.ID,userInfo.Id,loan,address,"update",key, pass)
	}
	if true == result {
		return ResponseFun("审核成功",200)
	}else {
		return ResponseFun("审核失败",30024)
	}
}


func IsBindingMiningAddressServices(loanMiningId,token string) Response {

	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",30028)
	}
	if !VerifyAdminRole(userInfo) {
		return ResponseFun("无权限操作",30030)
	}

	depositInfo := models.GetLoanMiningInfoById(loanMiningId)
	if 1 != depositInfo.State {
		return ResponseFun("操作错误",30032)
	}

	userDepositBalanceInfo := models.GetUserLoanMiningBalanceByEmail(depositInfo.Email)
	if "" == userDepositBalanceInfo.Address {
		return ResponseFun(false,200)
	}
	return ResponseFun(true,200)
}


func ExtractLoanMiningServices(token,password string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",30034)
	}

	if !CheckPassword(token,password) {
		return ResponseFun("密码错误",30035)
	}

	userDepositBalanceInfo := models.GetUserLoanMiningBalanceByEmail(userInfo.Email)
	if "" == userDepositBalanceInfo.Email ||  500000 > userDepositBalanceInfo.Loan {
		return ResponseFun("你没有开始挖矿",30036)
	}

	result := models.ExtractLoanMining(userDepositBalanceInfo.ID,userDepositBalanceInfo.Loan,userDepositBalanceInfo.Deposit,userDepositBalanceInfo.Email,userInfo.Id)
	if true != result {
		return ResponseFun("申请失败",30038)
	}

	return ResponseFun("申请成功",200)
}



func ExtractLoanMiningReviewServices(extractLoanMiningBalanceId,reason,token,password,statesStr string) Response {
	userInfo := GetUserInfoByToken(token)
	states,err:=strconv.Atoi(statesStr)
	if nil != err {
		return ResponseFun("参数错误",30040)
	}

	if 3 != states && 5 != states {
		return ResponseFun("参数错误",30042)
	}
	if "" == userInfo.Email {
		return ResponseFun("token 无效",30044)
	}
	if !VerifyAdminRole(userInfo) {
		return ResponseFun("无权限操作",30046)
	}
	if !CheckPassword(token,password) {
		return ResponseFun("密码错误",30048)
	}

	depositInfo := models.GetExtractLoanMiningBalanceInfoById(extractLoanMiningBalanceId)
	if 1 != depositInfo.State {
		return ResponseFun("操作错误",30050)
	}

	userDepositBalanceInfo := models.GetUserDepositBalanceByEmail(depositInfo.Email)

	var deposit,loan float64
	if 3 == states {
		deposit = depositInfo.Deposit + userDepositBalanceInfo.Balance
		//loan = depositInfo.Loan + userDepositBalanceInfo.Loan
	}
	if 5 == states {
		deposit = depositInfo.Deposit
		loan = depositInfo.Loan
	}

	result := models.UpdateExtractLoanMining(states,deposit,loan,reason,depositInfo.Email,depositInfo.ID,userInfo.Id)
	if true == result {
		return ResponseFun("审核成功",200)
	}else {
		return ResponseFun("审核失败",30052)
	}
}


type LoanMiningList struct {
	LoanMiningList []models.LoanMining   `json:"loan_mining_list"`
	Page	int	`json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

func GetLoanMiningListServices(pageStr,pageSizeStr,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",30054)
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

	return ResponseFun(LoanMiningList{
		LoanMiningList:models.GetLoanMiningListByEmail(userInfo.Email,page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.GetLoanMiningListCountByEmail(userInfo.Email),
	},200)
}


func AdminGetLoanMiningListServices(pageStr,pageSizeStr,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",30054)
	}
	if !VerifyAdminRole(userInfo) {
		return ResponseFun("无权限操作",20018)
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

	return ResponseFun(LoanMiningList{
		LoanMiningList:models.GetLoanMiningList(page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.GetLoanMiningListCount(),
	},200)
}


type ExtractLoanMiningList struct {
	ExtractLoanMiningList []models.ExtractLoanMiningBalance   `json:"loan_mining_list"`
	Page	int	`json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

func ExtractLoanMiningListServices(pageStr,pageSizeStr,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",30056)
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

	return ResponseFun(ExtractLoanMiningList{
		ExtractLoanMiningList:models.GetExtractLoanMiningBalanceListByEmail(userInfo.Email,page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.GetExtractLoanMiningBalanceListCountByEmail(userInfo.Email),
	},200)
}

func AdminExtractLoanMiningListServices(pageStr,pageSizeStr,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if !VerifyAdminRole(userInfo) {
		return ResponseFun("无权限操作",20018)
	}
	if "" == userInfo.Email {
		return ResponseFun("token 无效",30056)
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

	return ResponseFun(ExtractLoanMiningList{
		ExtractLoanMiningList:models.GetExtractLoanMiningBalanceList(page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.GetExtractLoanMiningBalanceListCount(),
	},200)
}


func UserLoanMiningBalanceServices(token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",30034)
	}
	userDepositBalanceInfo := models.GetUserLoanMiningBalanceByEmail(userInfo.Email)
	userDepositBalanceInfo.Key=""
	userDepositBalanceInfo.Password=""
	return ResponseFun(userDepositBalanceInfo,200)
}