package controller

import (
	"fmt"
	"time"
)

func BenefitCalculation()  {
	for true {
		time.Sleep(time.Second * 2)
		fmt.Println(time.Now())
	}
}


//计算收益
/**
* 1 根据小弟地址，获取余额，并转账道一个固定地址，转账金额，状态等信息入库。
* 2 等过五分钟，去检查交易状态
* 3 计算收益，个人收益=总收益*（挖矿保证金/总挖矿保证金）
* 4 入库，总收益 +=个人收益，可提现收益 += 个人收益
 */

func AutoTransaction()  {
	
}
 


 
