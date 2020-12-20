/**************************************
 * @Author: Yinjinlin
 * @Description:
 * @File:  reflectcase_test
 * @Version: 1.0.0
 * @Date: 2020/12/18 23:38
 **************************************/
package goreflect

import (
	"reflect"
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
	call1 := func(v1 int, v2 int) {
		t.Log(v1, v2)
	}
	call2 := func(v1, v2 int, s string) {
		t.Log(v1, v2, s)
	}
	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
	}
	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}

// 使用反射操作任意结构体
type User struct {
	UserId string
	Name   string
}

func TestReflectStruct(t *testing.T) {
	var (
		model *User
		sv    reflect.Value
	)
	model = &User{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf", sv.Kind().String())
	sv = sv.Elem()
	t.Log("reflect.ValueOf.Elem", sv.Kind().String())
	sv.FieldByName("UserId").SetString("12345678")
	sv.FieldByName("Name").SetString("jack")
	t.Log("model :", model)
}

// 使用反射创建并操作结构体
type Stud struct {
	UserId string
	Name   string
}

func TestReflectStructPtr(t *testing.T){
	var (
		model *Stud
		st reflect.Type
		elem reflect.Value
	)
	st = reflect.TypeOf(model)
	t.Log("reflect.TypeOf ", st.Kind().String()) // ptr
	st = st.Elem()
	t.Log("reflect.TypeOf.Elem",st.Kind().String())  // struct
	//New 返回一个value类型值，该值持有一个指向类型为type的新申请的零值的指针
	elem = reflect.New(st)
	t.Log("reflect.New", elem.Kind().String()) //ptr
	t.Log("reflect.New.Elem",elem.Elem().Kind().String())
	model = elem.Interface().(*Stud)
	elem = elem.Elem() //取得elem指向的值
	elem.FieldByName("UserId").SetString("shezhi123")
	elem.FieldByName("Name").SetString("jack-123")
	t.Log("model", model.Name, model.UserId)

}
