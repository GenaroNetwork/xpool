package controller

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestCreateUserServices(t *testing.T) {
	result := CreateUserServices("2581913653@qq.com","111111","FkJPh")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestGetVerificationCodeServices(t *testing.T) {
	result := GetVerificationCodeServices("2581913653@qq.com")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestLoginServices(t *testing.T) {
	result := LoginServices("2581913653@qq.com","111111")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestGetUserByTokenServices(t *testing.T) {
	result := GetUserByTokenServices("qC5TzVshcPiV2GjRZjHjx91salYCYXmFA3UyQEKVf3Y425v53v")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}