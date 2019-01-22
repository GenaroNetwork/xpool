package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"xpool/models"
	"strconv"
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


type TransactionInfo struct {
	From     string 	`json:"from"`
	To       string		`json:"to"`
	Value	 float64	`json:"value"`
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

	models.SaveDeposit(&models.Deposit{
		State:1,
		Email:userInfo.Email,
		Hash:hash,
		Value:transactionInfo.Value,
	})

	return ResponseFun("增加保证金成功",200)
}

type DepositList struct {
	DepositList []models.Deposit   `json:"depositList"`
	Page	int	`json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

func GetDepositListServices(pageStr,pageSizeStr,token string) Response {
	userInfo := GetUserInfoByToken(token)
	if "" == userInfo.Address {
		return ResponseFun("获取地址失败",20014)
	}
	page,err:=strconv.Atoi(pageStr)
	if nil != err {
		page = 1
	}
	pageSize,err:=strconv.Atoi(pageSizeStr)
	if nil != err {
		pageSize = 100
	}

	if 0 == page {
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