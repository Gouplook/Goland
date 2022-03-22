/**
 * @Author: yinjinlin
 * @File:  function_test
 * @Description:
 * @Date: 2021/5/27 下午3:52
 */

package gofunc

import (
	"fmt"
	"testing"
)

func TestGenPalyer(t *testing.T) {
	// 创建一个玩家生成器
	generator := GenPalyer("mosou")

	// 返回新创建玩家的姓名, 血量
	name, hp := generator()

	fmt.Println(name, hp)

}

// 斐波那契数列
func TestFib(t *testing.T) {
	f00 := Fib()
	fmt.Println(f00(), f00(), f00(), f00())
}

// 判断图片文件后缀,若没有加上jpg后缀(使用闭包）
func TestMakeSuffix(t *testing.T) {
	ff := MakeSuffix(".jpg")

	fmt.Println(ff("win"))
	fmt.Println(ff("jack.jpg"))

}

// 匿名函数
func TestAnonymous(t *testing.T) {
	Anonymous(20, 60)
}

func TestRecvertest(t *testing.T) {
	Recvertest()
}
