/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/23 下午3:25

*******************************************/
package tool

import (
	"strconv"
	"strings"
)

//字符串切割成int型数组
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

