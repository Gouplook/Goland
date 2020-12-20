/**************************************
 * @Author: Yinjinlin
 * @Description: 反射实践案例
 * @File:  reflectcase
 * @Version: 1.0.0
 * @Date: 2020/12/18 23:37
 ************************************/
package goreflect

import (
	"fmt"
	"reflect"
)

//定义了一个Monster结构体
type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

//方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//方法， 接收四个值，给s赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

//方法，显示s的值
func (s Monster) Print() {
	fmt.Println("---start~----")
	fmt.Println(s)
	fmt.Println("---end~----")
}

func StructCase(s interface{}) {
	rType := reflect.TypeOf(s)
	rVal := reflect.ValueOf(s)

	valKind := rVal.Kind()
	// 如果传入的不是struct，就退出
	if valKind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	// 获取到该结构体有几个字段
	structFieldNum := rVal.NumField()
	fmt.Printf("struct Field Num has %d fields\n", structFieldNum)
	// 遍历结构体的所有字段
	for i := 0; i < structFieldNum; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, rVal.Field(i))
		//获取到struct标签, 注意需要通过reflect.Type来获取tag标签的值
		tagVal := rType.Field(i).Tag.Get("json") // json 与结构体定义相同，可以自定义
		//如果该字段于tag标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := rVal.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
}

// --------------------------------



// ----------------------------------------------------------
// 编写一个Call结构，有两个字段Num1 Num2 方法 GetSub(name string)
// 使用反射遍历Call结构体所有的字段信息
// 使用反射机制完成对GetSub的调用，输出形式为：“Tom 完成了减法运算 8-5= 3”

