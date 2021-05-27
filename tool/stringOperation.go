/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/23 下午3:25
@Description:字符串问题转换总结

*******************************************/
package tool

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串切割成int型数组
// str := "112,34,56,78"  ---- >  [112 34 56 78]
func StrExplode2IntArr( s string, step string) []int {
	strs := strings.Split(s, ",")
	var outData []int
	for _, v := range strs{
		if len(v) == 0{
			continue
		}
		intv, _ :=strconv.Atoi(v)
		outData = append( outData, intv )
	}
	return outData
	// 1,2,5 类型卡 --> 适合这些门店 4 5 6 9
}

// TrimRgiht
func StringsTrim(s string, cutset string) string{
	 str := strings.TrimRight(s,cutset)
	 return str
}


//获取字符串长度
//@param  string str 待获取长度字符串
//@return int
func Mb4Strlen(str string) int{
	str = strings.TrimSpace(str)
	if len(str) == 0{
		return 0
	}
	strRune := []rune(str)
	lens := len(strRune)
	return lens
}

//截取字符串
//@param string str   待截取的字符串
//@param int    index 截取开始位置
//@param int    lens  截取长度
func StuffStr(str string,index int,lens int)(string){
	str = strings.TrimSpace(str)
	if len(str) == 0{
		return str
	}
	strRune := []rune(str)
	if len(strRune)<lens{
		lens = len(strRune)
	}
	return string(strRune[index:lens])

}

// 公钥转换
func GetPemPublic(public_key string) string {
	res := "-----BEGIN PUBLIC KEY-----\n"
	strlen := len(public_key)
	for i:=0;i < strlen;i+=64 {
		if i + 64 >= strlen {
			res += public_key[i:] + "\n"
		}else{
			res += public_key[i:i + 64] + "\n"
		}
	}
	res += "-----END PUBLIC KEY-----"
	return res
}

// 私钥转换
func GetPemPrivate(private_key string) string {
	res := "-----BEGIN RSA PRIVATE KEY-----\n"
	strlen := len(private_key)
	for i:=0;i < strlen;i+=64 {
		if i + 64 >= strlen {
			res += private_key[i:] + "\n"
		}else{
			res += private_key[i:i + 64] + "\n"
		}
	}
	res += "-----END RSA PRIVATE KEY-----"
	return res
}
//把数组转换为字符串
//@param  string separator       转换分隔符
//@param  interface{}  interface 待转换数据
//@return string
func Implode(separator string, array interface{}) (string) {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", separator, -1)
}