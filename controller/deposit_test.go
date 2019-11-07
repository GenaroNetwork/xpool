package controller

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAddDepositServices(t *testing.T) {
	result := AddDepositServices("0x46d91cbe2571c3cf32a663afbf24b89a8355d764b4069e97cc8bd74e9c4f1a07", "333333888855", "xHfZ6rb6IP5u8GD6FeTIe62rcwIPH7EVTizZAu8rXV6miTPwD3")
	test, _ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestGetDepositListServices(t *testing.T) {
	result := GetDepositListServices("1", "100", "xHfZ6rb6IP5u8GD6FeTIe62rcwIPH7EVTizZAu8rXV6miTPwD3")
	test, _ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestDepositReviewServices(t *testing.T) {
	result := DepositReviewServices("4", "okxxx", "3aGFEB1Imer3qL1fra2pWi6vST5zDjLesDFH0iIPy1kYDKTGNx", "123456", "3")

	//result := DepositReviewServices("4","okxxx","4S4MBytl3dydYGJK2oDlcgUjXUIYCV4Nqx4Q3Ye681znhkcnXs","333333888855","3")
	test, _ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestExtractDepositServices(t *testing.T) {
	result := ExtractDepositServices("pVdXGXLVbXsVqKZEEHiBmZ5Qmj2XEbEwDqOkZir8rVpSrOTvU3", "1", "333333888855")
	test, _ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestExtractDepositReviewServices(t *testing.T) {
	result := ExtractDepositReviewServices("2", "okxxx", "qqm8vUwwFjHOnDLkBBey0QxERPSH6AFxACtlPEuAF7J2BuNljS", "123456", "5")

	//result := DepositReviewServices("4","okxxx","4S4MBytl3dydYGJK2oDlcgUjXUIYCV4Nqx4Q3Ye681znhkcnXs","333333888855","3")
	test, _ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestGetExtractDepositListServices(t *testing.T) {
	result := GetExtractDepositListServices("1", "100", "pVdXGXLVbXsVqKZEEHiBmZ5Qmj2XEbEwDqOkZir8rVpSrOTvU3")
	test, _ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestAdminGetDepositListServices(t *testing.T) {
	result := AdminGetDepositListServices("1", "100", "h9cbpQvcz15zyrnx7Sn3qsfCfKzSTqLaU3foCYXPKRKOdgSz97")
	test, _ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestAdminGetExtractDepositListServices(t *testing.T) {
	result := AdminGetExtractDepositListServices("1", "100", "h9cbpQvcz15zyrnx7Sn3qsfCfKzSTqLaU3foCYXPKRKOdgSz97")
	test, _ := json.Marshal(result)
	fmt.Println(string(test[:]))
}
