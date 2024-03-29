/**
 * @Author: yinjinlin
 * @File:  goslice
 * @Description:
 * @Date: 2021/12/7 下午1:44
 */

package goslice

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"unsafe"
)

//slice 底层结构
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

//如果切片的容量小于 1024 个元素，于是扩容的时候就翻倍增加容量

func IntSort() {
	var a = [...]int{1, 8, 38, 2, 348, 484}
	//数组是值类型,不能直接排序，必须转为切片
	sort.Ints(a[:])
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// 议题
	// 1. 自由人管理，兼职人员管理
	//    兼职人员出图，合拼到一起，一定查图交给客户。
	//    有些外协人员，一定培养他们。
	// 2. 兼职人员，一套图纸，可以分单元外发。
	//    对于外协人员，首先要善意的提醒，第二保持他们的质量
	//    第三，尽量给他们提供正确且明确的参考及要求。
	// 3. 公司内部人员管理，最大程度减少人力额外成本。提高工作效率。
	// 4. 每天早上定一个时间，开例会，3件事，回顾昨天工作内容及完成进度
	//    总结昨天质量或进度存在的问题，布置今天的任务目标及完成计划（由项目负责人执行）。
	//
}

// 把数组转换为字符串
// @param  string separator       转换分隔符
// @param  interface{}  interface 待转换数据
// @return string

func Implode(separator string, array interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", separator, -1)
}

// slice 排序

func SliceSort() {
	//定义一个年龄列表
	ageList := []int{1, 3, 7, 7, 8, 2, 5}

	//排序，实现比较方法即可
	sort.Slice(ageList, func(i, j int) bool {
		return ageList[i] < ageList[j]
	})
	fmt.Printf("after sort:%v", ageList)
}

// 两个切片包含
// 判断 b中元素是否在a切片中，不在提示
// a= []int {1,2,4,9}
// b = []int {2,4}

func SliceContains(a, b []int) {
	//a := []int{0, 1, 3, 2}
	//
	//b := []int{1, 0}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	aa := make([]byte, 0)
	bb := make([]byte, 0)
	for _, v := range a {
		aa = append(aa, byte(v))
	}

	for _, v := range b {
		bb = append(bb, byte(v))
	}
	//fmt(a)
	fmt.Println(aa)

	fmt.Println(bb)
	fmt.Println(bytes.Contains(aa, bb))

}
