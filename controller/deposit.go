package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
	"xpool/models"
)

var Deposit deposit = deposit{}
type deposit struct{}


func (u *deposit) AddDeposit(c *gin.Context) {
	//money := c.PostForm("money")
	hash := c.PostForm("hash")
	password := c.PostForm("password")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,AddDepositServices(hash,password,token))
}


func (u *deposit) GetDepositList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,GetDepositListServices(page,pageSize,token))
}

func (u *deposit) AdminGetDepositList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,AdminGetDepositListServices(page,pageSize,token))
}

func (u *deposit) DepositReview(c *gin.Context) {
	depositId := c.PostForm("depositId")
	reason := c.PostForm("reason")
	token := c.PostForm("token")
	password := c.PostForm("password")
	states := c.PostForm("states")
	c.JSON(http.StatusOK,DepositReviewServices(depositId,reason,token,password,states))
}



func (u *deposit) ExtractDeposit(c *gin.Context) {
	value := c.PostForm("value")
	token := c.PostForm("token")
	password := c.PostForm("password")
	c.JSON(http.StatusOK,ExtractDepositServices(token,value,password))
}


func (u *deposit) ExtractDepositReview(c *gin.Context) {
	extractDepositId := c.PostForm("extractDepositId")
	reason := c.PostForm("reason")
	states := c.PostForm("states")
	token := c.PostForm("token")
	password := c.PostForm("password")
	c.JSON(http.StatusOK,ExtractDepositReviewServices(extractDepositId,reason,token,password,states))
}

func (u *deposit) GetExtractDepositList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,GetExtractDepositListServices(page,pageSize,token))
}

func (u *deposit) AdminGetExtractDepositList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.PostForm("pageSize")
	token := c.PostForm("token")
	c.JSON(http.StatusOK,AdminGetExtractDepositListServices(page,pageSize,token))
}

func (u *deposit) DepositBalance(c *gin.Context)  {
	token := c.PostForm("token")
	c.JSON(http.StatusOK,DepositBalanceServices(token))
}

type TransactionInfo struct {
	From     string 	`json:"from"`
	To       string		`json:"to"`
	Value	 string		`json:"value"`
	Status	 int		`json:"status"`
}


func AddDepositServices(hash,password,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Address {
		return ResponseFun("获取地址失败",20000)
	}

	if !CheckPassword(token,password) {
		return ResponseFun("密码错误",20002)
	}

	if 66 != len(hash) {
		return ResponseFun("hash 错误",20010)
	}

	result := HttpGet(TRANSACTIONINFO+hash)
	if nil == result {
		return ResponseFun("获取交易失败",20004)
	}
	var transactionInfo TransactionInfo
	err := json.Unmarshal(result, &transactionInfo)
	if nil != err {
		return ResponseFun("获取交易失败",20006)
	}

	if 1 != transactionInfo.Status || OFFICIALADDRESS != transactionInfo.To {
		return ResponseFun("获取交易失败",20008)
	}

	depositInfo := models.GetDepositInfoByHsah(hash)
	if "" != depositInfo.Hash {
		return ResponseFun("增加保证金Hash已存在",20012)
	}
	value := new(big.Int)
	value, _ = value.SetString(transactionInfo.Value,10)
	valueFloat,_ := new(big.Float).Quo(new(big.Float).SetInt(value), big.NewFloat(1000000000000000000)).Float64()

	if 0 == valueFloat {
		return ResponseFun("保证金为0",20011)
	}
	models.SaveDeposit(&models.Deposit{
		State:1,
		Email:userInfo.Email,
		Hash:hash,
		Value:valueFloat,
	})

	return ResponseFun("申请增加保证金成功",200)
}

type DepositList struct {
	DepositList []models.Deposit   `json:"depositList"`
	Page	int	`json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

func GetDepositListServices(pageStr,pageSizeStr,token string) Response {
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

	return ResponseFun(DepositList{
		DepositList:models.GetDepositListByEmail(userInfo.Email,page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.GetDepositCountByEmail(userInfo.Email),
	},200)
}

func AdminGetDepositListServices(pageStr,pageSizeStr,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if !VerifyAdminRole(userInfo) {
		return ResponseFun("无权限操作",20018)
	}
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

	return ResponseFun(DepositList{
		DepositList:models.GetDepositList(page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.GetDepositCount(),
	},200)
}

func DepositReviewServices(depositId,reason,token,password,statesStr string) Response {
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
	depositInfo := models.GetDepositInfoById(depositId)
	if 1 != depositInfo.State {
		return ResponseFun("操作错误",20022)
	}

	//检查是否是第一条数据
	userDepositBalanceInfo := models.GetUserDepositBalanceByEmail(depositInfo.Email)
	var value float64
	if 3 == states {
		value = depositInfo.Value + userDepositBalanceInfo.Balance
	}
	var result bool
	if "" == userDepositBalanceInfo.Email {
		result = models.UpdateDeposit(states,value,reason,depositInfo.Email,depositInfo.ID,userInfo.Id,"create")
	}else {
		result = models.UpdateDeposit(states,value,reason,depositInfo.Email,depositInfo.ID,userInfo.Id,"update")
	}
	if true == result {
		return ResponseFun("审核成功",200)
	}else {
		return ResponseFun("审核失败",20024)
	}

}


func ExtractDepositServices(token,valueStr,password string) Response {
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

	userDepositBalanceInfo := models.GetUserDepositBalanceByEmail(userInfo.Email)

	balance := Round(userDepositBalanceInfo.Balance - value,3)

	if 0 > balance {
		return ResponseFun("保证金余额不足",20032)
	}

	result := models.SaveExtractDeposit(1, userInfo.Email, value,balance,userInfo.Id)
	if true != result {
		return ResponseFun("申请提取保证金失败",20034)
	}

	return ResponseFun("申请提取保证金成功",200)
}



func ExtractDepositReviewServices(extractDepositId,reason,token,password,statesStr string) Response {
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
	depositInfo := models.GetExtractDepositInfoById(extractDepositId)
	if 1 != depositInfo.State {
		return ResponseFun("操作错误",20022)
	}

	userDepositBalanceInfo := models.GetUserDepositBalanceByEmail(depositInfo.Email)
	var value float64
	if 5 == states {
		value = depositInfo.Value + userDepositBalanceInfo.Balance
	}
	result := models.UpdateExtractDeposit(states,value,reason,depositInfo.Email,depositInfo.ID,userInfo.Id)
	if true == result {
		return ResponseFun("审核成功",200)
	}else {
		return ResponseFun("审核失败",20024)
	}

}


type ExtractDepositList struct {
	ExtractDepositList []models.ExtractDeposit   `json:"extract_deposit_list"`
	Page	int	`json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

func GetExtractDepositListServices(pageStr,pageSizeStr,token string) Response {
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

	return ResponseFun(ExtractDepositList{
		ExtractDepositList:models.GetExtractDepositListByEmail(userInfo.Email,page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.GetExtractDepositCountByEmail(userInfo.Email),
	},200)
}


func AdminGetExtractDepositListServices(pageStr,pageSizeStr,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",20014)
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

	return ResponseFun(ExtractDepositList{
		ExtractDepositList:models.GetExtractDepositList(page,pageSize),
		Page:page,
		PageSize:pageSize,
		Total:models.GetExtractDepositCount(),
	},200)
}

func DepositBalanceServices(token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Email {
		return ResponseFun("token 无效",20014)
	}
	return ResponseFun(Round(models.GetUserDepositBalanceByEmail(userInfo.Email).Balance,2),200)
}