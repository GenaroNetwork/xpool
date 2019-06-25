package autoTransaction

import (
	"testing"
	"fmt"
	"encoding/json"
	"math/big"
)

func TestUnlockAccount(t *testing.T) {
	//result := UnlockAccount("", AccountPassword)
	//var unlockAccountResult UnlockAccountResult
	//if nil == result {
	//
	//}
	//json.Unmarshal(result, &unlockAccountResult)
	//fmt.Println(unlockAccountResult.Result)
}


func TestGetBalance(t *testing.T) {
	result := GetBalance("0x11c14387b2ae26087f2700520361209ed2c7ab07")
	var getBalanceResult GetBalanceResult
	if nil == result {

	}
	err := json.Unmarshal(result, &getBalanceResult)
	if nil != err {

	}
	balance  := big.NewInt(0)
	fmt.Println(getBalanceResult.Result.ToInt())
	balance = balance.Sub(getBalanceResult.Result.ToInt(),big.NewInt(4139617255607405828))
	fmt.Println(balance)
	//SendTx("0x11c14387b2ae26087f2700520361209ed2c7ab07",balance)
	fmt.Println(getBalanceResult.Result.ToInt())
}

func TestSendtransaction(t *testing.T) {
	//Sendtransaction("0x525c4697c9db709793cd499bcdbcdffaf4565313")
}

func TestSendtransactionAll(t *testing.T)  {
	SendtransactionAll()
}