/**************************************
 * @Author: Yinjinlin
 * @Description:
 * @File:  reflectcase_test
 * @Version: 1.0.0
 * @Date: 2020/12/18 23:38
 **************************************/
package goreflect

import (
	"testing"
)

//反射实践案例
func TestStructCase(t *testing.T) {
	var a Monster = Monster{
		Name:  "jack ma",
		Age:   18,
		Score: 60.8,
	}
	StructCase(a)
}

// 定义一个适配器用作统一接口
func TestReflectFunc(t *testing.T) {
	ReflectFunc()
}

// 使用反射操作任意结构体
func TestReflectStruct(t *testing.T) {
	ReflectStruct()

}

// 使用反射创建并操作结构体
func TestReflectStructPtr(t *testing.T){
	ReflectStructPtr()
}

//
func TestCallReflect(t *testing.T) {
	model := &Call{}
	model.GetSub("tom")
	CallReflect(model)
}
