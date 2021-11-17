/**
 * @Author: yinjinlin
 * @File:  enumeration
 * @Description:
 * @Date: 2021/11/9 上午10:17
 */

package goenumeration

import "fmt"

// 枚举
type FishType int

const (
	A FishType= iota
	B
	C
	D
)

func (f FishType)String()string{
	return [...]string{
		"A",
		"B",
		"C",
		"D",
	}[f]
}


func Enmeration(){
	fmt.Println(A,B,C)
}