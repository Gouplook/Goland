/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/23 下午5:00

*******************************************/
package tool

import (
	"fmt"
	"testing"
)

func TestArrayKeys(t *testing.T) {
	ma := map[int]interface{}{
		1: "name",
		2: "age",
	}
	slile := ArrayKeys(ma)
	fmt.Println(slile)
}
func TestArrayString(t *testing.T) {
	tagIds := []string{"1#","2#","#4"}
	fmt.Println("tagIds: ",tagIds)
	str := ArrayString(",", tagIds)
	fmt.Println(str)
}
