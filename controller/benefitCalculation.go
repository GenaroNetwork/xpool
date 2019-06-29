package controller

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"time"
	"xpool/autoTransaction"
	"xpool/models"
)

func BenefitCalculation()  {
	for true {
		time.Sleep(time.Second * 2)
		fmt.Println(time.Now())
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
		result,hash := AutoTransaction(v.Email)
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

func AutoTransaction(email string) (*big.Int,string) {
	result := models.GetUserLoanMiningBalanceByEmail(email)
	d1 := []byte(result.Key)
	err := ioutil.WriteFile("./keystore", d1, 0666)
	if nil != err {
		fmt.Println(err.Error())
	}
	return autoTransaction.Sendtransaction(result.Address,BENEFITCALCULATION,
		"./keystore","123456")
}