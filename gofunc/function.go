/**
 * @Author: yinjinlin
 * @File:  function
 * @Description:
 * @Date: 2021/5/27 下午3:52
 */

package gofunc

import (
	"fmt"
	"strings"
)

// 函数 + 引用外部变量 = 闭包

// 函数闭包
func Square(x int) int {
	return x * x
}

func GenPalyer(name string) func() (string, int) {
	// 定义玩家血量
	hp := 1000
	return func() (string, int) {
		return name, hp
	}
}

// 斐波那契数列
func Fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 1)编写一个函数 makeSuffix(suffix string)  可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg ,
//   如果已经有.jpg后缀，则返回原文件名。
// 3)要求使用闭包的方式完成
// 4)strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。
func MakeSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 匿名函数
func Anonymous(a int, b int) {

	res := func(a int, b int) int {
		return a + b
	}
	rs := res(a, b)

	fmt.Println(rs)
}

//捕捉异常
func Recvertest() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	sum := num1 / num2

	fmt.Println(sum)
}
