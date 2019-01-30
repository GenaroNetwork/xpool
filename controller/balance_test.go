package controller

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestExtractBalanceServices(t *testing.T) {
	result := ExtractBalanceServices("OCSIwRWpbTpnpoA21kSJKlG5hAXLEnJY0wxwdfWDTd9ZS89zIT","123456")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestAddUserBalanceServices(t *testing.T)  {
	AddUserBalanceServices(524568.55,"2581913653@qq.com")
}

func TestExtractBalanceReviewServices(t *testing.T) {
	result := ExtractBalanceReviewServices("4","","y51LGgLsMj3TBor9dkoZv1IhSI2LEnc5Z0vnZhDub01Zq7pRnW","123456","5")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}