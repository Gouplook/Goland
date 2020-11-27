/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/23 下午3:27

*******************************************/
package tool

import (
	"fmt"
	"testing"
)

func TestStrExplode2IntArr(t *testing.T) {
	str := "112,34,56,78* 88  990"
	rs := StrExplode2IntArr(str,",")
	fmt.Println(rs)
}
