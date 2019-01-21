package controller

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
	"regexp"
)

type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}


/*
 * 封装返回的结果
 * status＝0 表示成功，status＝1表示未登录，status＝3 参数操作
 */
func ResponseFun(data interface{},status int) Response {
	result := Response{}
	result.Status = status
	result.Data = data
	return result
}



// 生成32位MD5
func MD5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}



//生成随机字符串
func GetRandomString(lengeth int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lengeth; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
