/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/28 09:28
@Description:

*********************************************/
package gostruct

import (
	"fmt"
)

type Demo struct {
	Name  string
	Ptr   *string
	Slice []int
	maps  map[string]string
}
// 结构体中基本应用
func StructBase() {
	// 说明：slice 不需要make，map和指针 使用之前必须make
	var d1 Demo
	d1.Name = "Aollo"
	d1.Slice = []int{1,2}

	d1.maps = make(map[string]string)
	d1.maps["CradId"] = "Uid"

	d1.Ptr = new(string)
    *(d1.Ptr) = "ptrString"

	fmt.Println(d1)
	fmt.Println(*(d1.Ptr))
}
