package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"time"
	"xpool/autoTransaction"
	"xpool/models"
)

func BenefitCalculation()  {
	for true {
		time.Sleep(time.Second * 7200)
		UserLoanMiningBalance()
		time.Sleep(time.Second * 300)
		ChechTransactionInfo()
		time.Sleep(time.Second * 300)
		Calculation()
	}
}

func BalanceFloat(result *big.Int) float64 {
	oneGnx := big.NewInt(1000000000000000000)

	float0point3Gnx := big.NewInt(1000000000000000)
	//获取整数部分
	resultTmp := big.NewInt(0)
	resultTmp.Div(result,oneGnx)
	//获取小数部分
	resultTmp2 := big.NewInt(0)
	resultTmp2.Mul(resultTmp,oneGnx)
	//获取小数
	resultTmp3 := big.NewInt(0)
	resultTmp3.Sub(result,resultTmp2)
	//转化小数
	resultTmp4 := big.NewInt(0)
	resultTmp4.Div(resultTmp3,float0point3Gnx)
	//计算总和
	return float64(resultTmp.Int64())+float64(resultTmp4.Int64())/1000.0
}

func UserLoanMiningBalance() {
	userLoanMining := models.GetUserLoanMiningBalance()
	for _,v:= range userLoanMining {
		result,hash := AutoTransaction(v.Email,v.Password)
		balanceFloat :=  BalanceFloat(result)
		if 0 == balanceFloat {
			continue
		}
		err := models.SaveAutoTransaction(&models.AutoTransaction{
			Email:v.Email,
			From:v.Address,
			To:BENEFITCALCULATION,
			Value:balanceFloat,
			Hash:hash,
			Tag:0,
			Status:0,
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}

//计算收益
/**
* 1 根据小弟地址，获取余额，并转账道一个固定地址，转账金额，状态等信息入库。
* 2 等过五分钟，去检查交易状态
* 3 计算收益，个人收益=总收益*（挖矿保证金/总挖矿保证金）
* 4 入库，总收益 +=个人收益，可提现收益 += 个人收益
 */

func AutoTransaction(email,pass string) (*big.Int,string) {
	result := models.GetUserLoanMiningBalanceByEmail(email)
	d1 := []byte(result.Key)
	err := ioutil.WriteFile("./keystore", d1, 0666)
	if nil != err {
		fmt.Println(err.Error())
	}
	return autoTransaction.Sendtransaction(result.Address,BENEFITCALCULATION,
		"./keystore",pass)
}

func ChechTransactionInfo()  {
	result := models.GetAutoTransaction(0)
	for _,v := range result {
		chechResult :=  GetTransactionInfo(v.Hash)
		fmt.Println(chechResult)
		if true == chechResult {
			models.UpdateAutoTransaction(v.Hash)
		}
	}
}


func GetTransactionInfo(hash string) bool {
	result := HttpGet(TRANSACTIONINFO+hash)
	if nil == result {
		fmt.Println("获取交易失败")
		return false
	}
	var transactionInfo TransactionInfo
	err := json.Unmarshal(result, &transactionInfo)
	if nil != err {
		fmt.Println(string(result[:]))
		fmt.Println("获取交易失败"+err.Error())
		return false
	}

	if 1 != transactionInfo.Status || BENEFITCALCULATION != transactionInfo.To {
		return false
	}
	return true
}

//计算
func Calculation()  {
	//获取当前总收益
	result := models.GetAutoTransaction(1)
	resultValtTotal := 0.0
	for _,v := range result {
		resultValtTotal += v.Value
	}
	if 0 == resultValtTotal {
		return
	}
	//获取矿池挖矿资金
	depositValtTotal := 0.0
	userLoanMining := models.GetUserLoanMiningBalance()
	for _,v := range userLoanMining {
		depositValtTotal += v.Deposit
	}
	if 0 == depositValtTotal {
		return
	}
	fmt.Println("########11")
	for _,v := range userLoanMining {
		userIncome := resultValtTotal*v.Deposit/depositValtTotal*(1-HANDLINGFEE)
		userIncome = Decimal(userIncome)
		fmt.Println(userIncome)
		getIncomeInfoById := models.GetIncomeInfoById(v.Email)
		if 0 == getIncomeInfoById.ID {
			models.UpdateIncome(v.Email,userIncome,userIncome,0,"create")
		}else {
			models.UpdateIncome(v.Email,getIncomeInfoById.TotalIncome+userIncome,getIncomeInfoById.IncomeBalance+userIncome,0,"update")
		}
	}
	models.UpdateAutoTransactionTag(1)
	//fmt.Println(resultValtTotal)
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}