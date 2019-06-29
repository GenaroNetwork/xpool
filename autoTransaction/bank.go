package autoTransaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GenaroNetwork/GenaroCore/common/hexutil"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
)

type UnlockAccountParameter struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      int      `json:"id"`
}

type UnlockAccountResult struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  bool `json:"result"`
}


type GetBalanceParameter struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      int      `json:"id"`
}

type GetBalanceResult struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  *hexutil.Big `json:"result"`
}



type GetHeftListResult struct {
	Result  map[string]*hexutil.Big `json:"result"`
}

type SendTxArgs struct {
	From      string         `json:"from"`
	To        string         `json:"to"`
	Value     *big.Int         `json:"value"`
}

type SendTxParameter struct {
	Jsonrpc string       `json:"jsonrpc"`
	Method  string       `json:"method"`
	Params  []SendTxArgs `json:"params"`
	Id      int          `json:"id"`
}

type send_transaction struct {
	Jsonrpc string       `json:"jsonrpc"`
	Method  string       `json:"method"`
	Params  []string 		 `json:"params"`
	Id      int          `json:"id"`
}


func UnlockAccount(account, password string) []byte {
	if "" == account || "" == password {
		return nil
	}
	parameter := UnlockAccountParameter{
		Jsonrpc: "2.0",
		Method:  "personal_unlockAccount",
		Id:      1,
	}
	parameter.Params = append(parameter.Params, account)
	parameter.Params = append(parameter.Params, password)
	input, _ := json.Marshal(parameter)
	result := httpPost(input)
	if nil == result {
		return nil
	}
	return result
}

func GetBalance(account string) []byte {
	if "" == account  {
		return nil
	}
	parameter := GetBalanceParameter{
		Jsonrpc: "2.0",
		Method:  "eth_getBalance",
		Id:      1,
	}
	parameter.Params = append(parameter.Params, account)
	parameter.Params = append(parameter.Params,"latest")
	input, _ := json.Marshal(parameter)
	result := httpPost(input)
	if nil == result {
		return nil
	}
	return result
}


func SendRawTransaction(sendTxArgs string) (bool,[]byte) {
	parameter := send_transaction{
		Jsonrpc: "2.0",
		Method:  "eth_sendRawTransaction",
		Id:      1,
	}
	parameter.Params = append(parameter.Params, sendTxArgs)
	input, _ := json.Marshal(parameter)
	fmt.Println(string(input[:]))
	result := httpPost(input)
	if nil == result {
		return false,[]byte("")
	}
	fmt.Println(string(result[:]))
	return true,result
}

//func SendtransactionAll()  {
//	time.Sleep(time.Duration(RandInt64(100,300))*time.Second)
//	Sendtransaction("0x75b6acf75064674dbd1ee275a85cda93f7d6dd92","0xc0ffc7800ce9c9ad27f89999748d938908fe066f")
//	time.Sleep(time.Duration(RandInt64(100,300))*time.Second)
//	Sendtransaction("0xc0ffc7800ce9c9ad27f89999748d938908fe066f","0xfd0a558fcfe003f055e43cfefa57830f50d1761d")
//	time.Sleep(time.Duration(RandInt64(100,300))*time.Second)
//	Sendtransaction("0xfd0a558fcfe003f055e43cfefa57830f50d1761d","0x1eb5c9a661856c19b0591a0bec8b42ced70e478e")
//	time.Sleep(time.Duration(RandInt64(100,300))*time.Second)
//	Sendtransaction("0x1eb5c9a661856c19b0591a0bec8b42ced70e478e","0x75b6acf75064674dbd1ee275a85cda93f7d6dd92")
//}


func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

type RawTransactionResult struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

func Sendtransaction(account,To,keyDir,AccountPassword string) (*big.Int,string) {
	result := GetBalance(account)
	var getBalanceResult GetBalanceResult
	if nil == result {

	}
	err := json.Unmarshal(result, &getBalanceResult)
	if nil != err {

	}
	balance  := big.NewInt(0)
	balance = balance.Sub(getBalanceResult.Result.ToInt(),SubGas)
	if balance.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0),""
	}
	balance = big.NewInt(10000000000000000)
	Nonce := GetTransactionCount(account)
	transaction := RawTransaction(keyDir,AccountPassword,balance,Nonce.Uint64(),To)
	if "" == transaction {
		return big.NewInt(0),""
	}

	resultBool, resultRawTra := SendRawTransaction(transaction)
	if false ==resultBool {
		return big.NewInt(0),""
	}

	var rawTransactionResult RawTransactionResult
	err = json.Unmarshal(resultRawTra,&rawTransactionResult)
	if err != nil {
		return big.NewInt(0),""
	}
	return balance,rawTransactionResult.Result
}



func httpPost(parameter []byte) []byte {
	if nil == parameter {
		return nil
	}
	client := &http.Client{}
	req_parameter := bytes.NewBuffer(parameter)
	request, _ := http.NewRequest("POST", ServeUrl, req_parameter)
	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	if nil == err && response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		return body
	}
	return nil
}



func GetTransactionCount(address string) *big.Int {
	if "" == address {
		return nil
	}
	parameter := GetBalanceParameter{
		Jsonrpc: "2.0",
		Method:  "eth_getTransactionCount",
		Id:      1,
	}
	parameter.Params = append(parameter.Params,address)
	parameter.Params = append(parameter.Params,"pending")
	input, _ := json.Marshal(parameter)
	result := httpPost(input)
	if nil == result {
		return nil
	}

	var getTransactionCount GetBalanceResult
	if nil == result {
		return nil
	}
	err := json.Unmarshal(result, &getTransactionCount)
	if nil != err {
		return nil
	}
	if nil == getTransactionCount.Result.ToInt() || getTransactionCount.Result.ToInt().Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	return getTransactionCount.Result.ToInt()
}