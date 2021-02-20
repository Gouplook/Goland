 /******************************************
 * @Author: Yinjinlin
 * @Description: 反射基本点
 * @File:  reflect_test
 * @Version: 1.0.0
 * @Date: 2020/12/18 21:48
 *****************************************/
package goreflect

import (
	"fmt"
	"testing"
)

// 基本类型反射测试
func TestBaseType(t *testing.T) {
	var num int = 10
	BaseType(num)
}


// 结构体反射
func TestStructType(t *testing.T) {
	stu := Student{
		Name: "jack",
		Age: 18,
	}
	StructType(stu)
}

// 通过反射，修改,
// num int 的值
// 修改 student的值
func TestReflectModeValue(t *testing.T) {
	// 原理：
	// num := 9
	// ptr *int = &num
	// num2 := *ptr  // === 类似 rVal.Elem()
	var num int = 10
	ReflectModeValue(&num)
	fmt.Println(num)
}
