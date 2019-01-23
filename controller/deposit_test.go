package controller

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestAddDepositServices(t *testing.T) {
	result := AddDepositServices("0x46d91cbe2571c3cf32a663afbf24b89a8355d764b4069e97cc8bd74e9c4f1a07","333333888855","xHfZ6rb6IP5u8GD6FeTIe62rcwIPH7EVTizZAu8rXV6miTPwD3")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestGetDepositListServices(t *testing.T)  {
	result := GetDepositListServices("1","100","xHfZ6rb6IP5u8GD6FeTIe62rcwIPH7EVTizZAu8rXV6miTPwD3")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestDepositReviewServices(t *testing.T)  {
	result := DepositReviewServices("4","okxxx","3aGFEB1Imer3qL1fra2pWi6vST5zDjLesDFH0iIPy1kYDKTGNx","123456","3")

	//result := DepositReviewServices("4","okxxx","4S4MBytl3dydYGJK2oDlcgUjXUIYCV4Nqx4Q3Ye681znhkcnXs","333333888855","3")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}