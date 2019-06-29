package models

import "testing"

func TestUpdateUserLoanMiningBalance(t *testing.T) {
	key := `{"address":"e25de09fb1afd3cacfad2e91cf5d5f2862597667","crypto":{"cipher":"aes-128-ctr","ciphertext":"f5c14d33b2909ed1053fe78921dba6bf20e971e67904bfcd2f26abd3170be2a1","cipherparams":{"iv":"4553eaaba93fc11d50153d61b9e4b829"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"59a4f1f81abd3e53409b7e06ffc147680df4d6f53f15244ce4ae6a7934198778"},"mac":"694f982e90921d4791a6b193814f55754163a88ba6799661c5cb5bcdbd1e175d"},"id":"6cf3233b-0ab7-46d9-97a0-af60c38fc5eb","version":3}`
	UpdateUserLoanMiningBalance("1065482100@qq.com","123456",key)
}