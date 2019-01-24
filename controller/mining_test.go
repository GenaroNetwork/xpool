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
	result := LoanMiningReviewServices("3","okxxx","qqm8vUwwFjHOnDLkBBey0QxERPSH6AFxACtlPEuAF7J2BuNljS","123456","3")

	//result := LoanMiningReviewServices("4","okxxx","4S4MBytl3dydYGJK2oDlcgUjXUIYCV4Nqx4Q3Ye681znhkcnXs","333333888855","3")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}