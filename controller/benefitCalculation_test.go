package controller

import (
	"fmt"
	"math/big"
	"testing"
)

func TestAutoTransaction(t *testing.T) {
	oneGnx := big.NewInt(1000000000000000000)

	float0point3Gnx := big.NewInt(1000000000000000)

	result, hash := AutoTransaction("1065482100@qq.com", "123456")
	//获取整数部分
	resultTmp := big.NewInt(0)
	resultTmp.Div(result, oneGnx)
	//获取小数部分
	resultTmp2 := big.NewInt(0)
	resultTmp2.Mul(resultTmp, oneGnx)
	//获取小数
	resultTmp3 := big.NewInt(0)
	resultTmp3.Sub(result, resultTmp2)
	//转化小数
	resultTmp4 := big.NewInt(0)
	resultTmp4.Div(resultTmp3, float0point3Gnx)
	//计算总和
	resultTmp5 := float64(resultTmp.Int64()) + float64(resultTmp4.Int64())/1000.0
	fmt.Println(resultTmp5)
	fmt.Println(hash)
}

func TestUserLoanMiningBalance(t *testing.T) {
	UserLoanMiningBalance()
	ChechTransactionInfo()
	Calculation()
}
