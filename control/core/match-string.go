package core

import (
	"crypto/sha1"
	"encoding/hex"
	"regexp"
	"time"
	"math/rand"
	"strconv"
	"strings"
)

//Authentication and query modules
type MatchString struct {
}

//检查用户名
//param str string 用户名
//param min int 最少
//param max int 最长
//return bool 是否正确
func (this *MatchString) CheckUsername(str string,min int,max int) bool {
	return this.matchStr("^[a-zA-Z][a-zA-Z0-9_]{"+strconv.Itoa(min)+","+strconv.Itoa(max)+"}$", str)
}

//验证邮箱
//param str string 邮箱地址
//return bool 是否正确
func (this *MatchString) CheckEmail(str string) bool {
	return this.matchStr("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+$", str)
}

//验证密码
//param str string 密码
//param min int 最少
//param max int 最长
//return bool 是否正确
func (this *MatchString) CheckPassword(str string,min int,max int) bool {
	return this.matchStr("^[a-zA-Z0-9]{"+strconv.Itoa(min)+","+strconv.Itoa(max)+",20}$", str)
}

//获取字符串的SHA1值
//param content string 要计算的字符串
//return string 计算出的SHA1值
func (this *MatchString) GetSha1(content string) string {
	hasher := sha1.New()
	_, err = hasher.Write([]byte(content))
	if err != nil {
		SendLog(err.Error())
		return ""
	}
	sha := hasher.Sum(nil)
	return hex.EncodeToString(sha)
}

//匹配验证
//param str string 要验证的字符串
//param mStr string 验证
//return bool 是否成功
func (this *MatchString) matchStr(str string, mStr string) bool {
	res, err := regexp.MatchString(mStr, str)
	if err != nil{
		SendLog(err.Error())
		return false
	}
	return res == true
}

//获取随机字符串
//param n int 随机码
//return string 新随机字符串
func (this *MatchString) GetRandStr(n int) string{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	re := r.Intn(n)
	return strconv.Itoa(re)
}

//截取字符串
//param str string 要截取的字符串
//param star int 开始位置
//param length int 长度
//return string 新字符串
func (this *MatchString) SubStr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

//分解URL获取名称和类型
//param sendURL URL地址
//return map[string]string 返回值集合
func (this *MatchString) GetURLNameType(sendURL string) map[string]string {
	res := map[string]string{
		"full-name": "",
		"only-name": "",
		"type": "",
	}
	urls := strings.Split(sendURL, "/")
	if len(urls) < 1 {
		return res
	}
	res["full-name"] = urls[len(urls)-1]
	if res["full-name"] == "" {
		res["only-name"] = res["full-name"]
		return res
	}
	names := strings.Split(res["full-name"], ".")
	if len(names) < 2 {
		return res
	}
	res["type"] = names[len(names)-1]
	for i := 0 ; i <= len(names) ; i ++{
		if i == len(names) - 1{
			break
		}
		if res["only-name"] == ""{
			res["only-name"] = names[i]
		}else{
			res["only-name"] += "." + names[i]
		}
	}
	return res
}