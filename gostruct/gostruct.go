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
func sStructBase() {
	// 说明：slice 不需要make，map和指针 使用之前必须make和new
	var d1 Demo
	d1.Name = "Aollo"
	d1.Slice = []int{1, 2}

	d1.maps = make(map[string]string)
	d1.maps["CradId"] = "Uid"

	d1.Ptr = new(string)
	*(d1.Ptr) = "ptrString"

	fmt.Println(d1)
	fmt.Println(*(d1.Ptr))
	fmt.Println("Aollo")
	fmt.Println(d1.maps)
	var A, B int
	var C float64
	A = 2415 + 29490 + 2712 + 17205
	B = (82000 + 81000) * 0.6
	C = 10336 + 12835

	fmt.Println("A=", A)
	fmt.Println("B=", B)
	fmt.Println("C=", C)
}

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func sStructAndSlice() {
	var s []Student = []Student{
		{
			101, "zhangsan", 18, "beijing",
		},
		{
			102, "lisi", 28, "beijing",
		},
	}

	fmt.Println(s)

}

// 结构体与map
func StructAndMap() {
	m := make(map[int]Student)
	m[1] = Student{
		1011,
		"wang",
		20,
		"shanghai",
	}
	fmt.Println(m[1])
	fmt.Println(m[1].name)

}
