package controller

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestLoanMiningServices(t *testing.T) {
	result := LoanMiningServices("pVdXGXLVbXsVqKZEEHiBmZ5Qmj2XEbEwDqOkZir8rVpSrOTvU3","6","333333888855")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestLoanMiningReviewServices(t *testing.T) {
	result := LoanMiningReviewServices("9","okxxx","cUbHeYDx2pGKKDe9UM3bgianaIJcMRSpqF2bnFMtwMYDf9bZxR","123456","3","","","")

	//result := LoanMiningReviewServices("4","okxxx","4S4MBytl3dydYGJK2oDlcgUjXUIYCV4Nqx4Q3Ye681znhkcnXs","333333888855","3")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestExtractLoanMiningServices(t *testing.T) {
	result:= ExtractLoanMiningServices("ctJnGYQ6lVcl4sAXG7T5y0ltlqlGC1PYFCdy1AVYUx0PelBH66","333333888855")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestExtractLoanMiningReviewServices(t *testing.T) {
	result := ExtractLoanMiningReviewServices("9","ok","cUbHeYDx2pGKKDe9UM3bgianaIJcMRSpqF2bnFMtwMYDf9bZxR","123456","5")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestAdminGetLoanMiningListServices(t *testing.T) {
	result := AdminGetLoanMiningListServices("1","100","h9cbpQvcz15zyrnx7Sn3qsfCfKzSTqLaU3foCYXPKRKOdgSz97")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestAdminExtractLoanMiningListServices(t *testing.T) {
	result := AdminExtractLoanMiningListServices("1","100","h9cbpQvcz15zyrnx7Sn3qsfCfKzSTqLaU3foCYXPKRKOdgSz97")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}