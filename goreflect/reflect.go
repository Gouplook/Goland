/*********************************************************************
 * @Author: Yinjinlin
 * @Description: 基本数据类型、interface{}、reflect.Value)进行反射的基本操作
 * @File:  reflect
 * @Version: 1.0.0
 * @Date: 2020/12/18 21:38
 **********************************************************************/
package goreflect

import (
	"fmt"
	"reflect"
)

// 基本类型反射
func BaseType(b interface{}) {
	// int(float64)  ----> interface
	// 通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)

	// 普通类型kind相同
	typeKind := rType.Kind()
	valKind := rVal.Kind()
	fmt.Printf("typeKind = %v   valKind = %v ", typeKind, valKind)
	fmt.Println()

	// 获得rVal值,其实rVal值不能进行运算
	num := rVal.Int()
	fmt.Println(num)
	fmt.Printf("rVal= %v rVal type = %T", rVal, rVal) //rVal= 10 rVal type = reflect.Value

	// rVal --> int
	// reflect.Value--> interface ---> int
	iv := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	numInt := iv.(int)
	fmt.Println(numInt)
}

type Student struct {
	Name string
	Age int
}
// 结构体反射
func StructType(s interface{}) {
	// struct ---> interface
	// 1: 先获得reflect.Type
	rType := reflect.TypeOf(s)
	fmt.Println(rType)

	//2: 获得reflect.value
	rVal := reflect.ValueOf(s)
	fmt.Println(rVal)
	// 结构体kind可能相同，也可能不同，有包前缀
	//3:获取对应的变量kind（类别）有以下两种方式
	//(1) rVal.Kind() ==>
	valKind := rVal.Kind()
	//(2) rTyp.Kind()
	typKind := rType.Kind()
	// 结构体序列号


	fmt.Printf("valKind = %v   typKind = %v ", valKind, typKind)
	fmt.Println()

	// 反射是运行时的反射，只有通过类型断言取变量的值
	// rVal --> struct
	// reflect.Value--> interface ---> int
	iv := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name = %v",stu.Name)
	}
}

// 反射修改值
func ReflectModeValue(r interface{}){
	// 1:获取到 reflect.Value
	rVal := reflect.ValueOf(r)
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	//2. rVal Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
	rVal.Elem().SetInt(20)

	// 总结：
}
