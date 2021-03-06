package controller

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"time"
	"xpool/models"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

/*
 * 封装返回的结果
 * status＝0 表示成功，status＝1表示未登录，status＝3 参数操作
 */
func ResponseFun(data interface{}, code int) Response {
	result := Response{}
	result.Code = code
	result.Data = data
	return result
}

// 生成32位MD5
func MD5(text string) string {
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

func VerifyEthAdderss(adderss string) bool {
	pattern := `^(0x)?[0-9a-fA-F]{40}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(adderss)
}

func HttpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return body
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func CheckPassword(token, password string) bool {
	UserInfo := GetUserInfoByToken(token)
	if "" == UserInfo.Email {
		return false
	}
	getUser := models.GetUserByEmail(UserInfo.Email)
	passwordVal := MD5(getUser.SaltValue + MD5(password+getUser.SaltValue) + getUser.SaltValue)
	if passwordVal != getUser.Password {
		return false
	}
	return true
}

func VerifyAdminRole(userInfo UserInfo) bool {
	if 1 == userInfo.Role {
		return true
	}
	return false
}
