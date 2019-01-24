package controller

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestCreateUserServices(t *testing.T) {
	result := CreateUserServices("1065482100@qq.com","123456","g4f6C","0x572856549d51f68ebcc8f15a2749d65874131a29")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestGetVerificationCodeServices(t *testing.T) {
	result := GetVerificationCodeServices("1065482100@qq.com")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestLoginServices(t *testing.T) {
	//result := LoginServices("2581913653@qq.com","333333888855")
	result := LoginServices("1065482100@qq.com","123456")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestGetUserByTokenServices(t *testing.T) {
	result := GetUserByTokenServices("dikcggoeqBdELKIL08I3nS5TrpMcrF3OyPMumM5vsn70JgJBqs")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestForgetPasswordServices(t *testing.T) {
	result := ForgetPasswordServices("2581913653@qq.com","33333333","G4iYL")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}

func TestDeleteVerificationCode(t *testing.T) {
	DeleteVerificationCode("kTg0x")
}

func TestResetPasswordServices(t *testing.T) {
	result :=ResetPasswordServices("LfUyyeD08VFsXs3YRPFkLSRPk6nsU19HHd0R7yLCggq75pbcDi","333333888855","333333888855")
	test,_ := json.Marshal(result)
	fmt.Println(string(test[:]))
}