/**
 * @Author: yinjinlin
 * @File:  goslice
 * @Description:
 * @Date: 2021/12/7 下午1:44
 */

package goslice

import (
	"fmt"
	"sort"
)

func IntSort() {
	var a = [...]int{1, 8, 38, 2, 348, 484}
	//数组是值类型,不能直接排序，必须转为切片
	sort.Ints(a[:])
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
}
