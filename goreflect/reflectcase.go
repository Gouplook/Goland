/**************************************
 * @Author: Yinjinlin
 * @Description: 反射实践案例,
 * 总结：涉及指针问题需要Elem（）函数调用
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
	fmt.Println("---end~------")
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

// ----------------------------定义一个适配器用作统一接口-----------------------------
func ReflectFunc() {
	call1 := func(v1 int, v2 int) {
		fmt.Println(v1, v2)
	}
	call2 := func(v1, v2 int, s string) {
		fmt.Println(v1, v2, s)
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

// -------------------------使用反射操作任意结构体---------------------------------
// 使用反射操作任意结构体
type User struct {
	UserId string
	Name   string
}

func ReflectStruct() {
	var (
		model *User
		sv    reflect.Value
	)
	model = &User{}
	sv = reflect.ValueOf(model)
	numField :=  sv.Elem().NumField()
	fmt.Println(numField)

	fmt.Println("reflect.ValueOf", sv.Kind().String())
	sv = sv.Elem()
	fmt.Println("reflect.ValueOf.Elem", sv.Kind().String())
	sv.FieldByName("UserId").SetString("12345678")
	sv.FieldByName("Name").SetString("jack")
	fmt.Println("model :", model)
}

// --------------------------使用反射创建并操作结构体--------------------------------
type Stud struct {
	UserId string
	Name   string
}

func ReflectStructPtr() {
	var (
		model *Stud
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)
	fmt.Println("reflect.TypeOf ", st.Kind().String())     // ptr
	st = st.Elem()                                         // 若st的kind不是interface或ptr，会panic，
	fmt.Println("reflect.TypeOf.Elem", st.Kind().String()) // struct
	//New 返回一个value类型值，该值持有一个指向类型为type的新申请的零值的指针
	elem = reflect.New(st)
	fmt.Println("reflect.New", elem.Kind().String())             //ptr
	fmt.Println("reflect.New.Elem", elem.Elem().Kind().String()) // struct
	// 返回elem当前持有的值，表示为/保管在interface了{}类型
	model = elem.Interface().(*Stud)
	elem = elem.Elem() //取得elem指向的值
	elem.FieldByName("UserId").SetString("shezhi123")
	elem.FieldByName("Name").SetString("jack-123")
	fmt.Println("model", model.Name, model.UserId)

	// 创建对象 需要申请空间
	elem = reflect.New(st)
	createObj := elem.Interface().(*Stud)
	elem = elem.Elem()
	elem.FieldByName("UserId").SetString("0001")
	elem.FieldByName("Name").SetString("Jack-createObj")

	fmt.Println("createObj", createObj.UserId, createObj.Name)
}

// -------------------------------案例-----------------------------
// 编写一个Call结构，有两个字段Num1 Num2 方法 GetSub(name string)
// 使用反射遍历Call结构体所有的字段信息
// 使用反射机制完成对GetSub的调用，输出形式为：“Tom 完成了减法运算 8-5= 3”

type Call struct {
	Name string `json:"name"`
	Num1 int	`json:"num_1"`
	Num2 int	`json:"num_2"`
}

func (c *Call) GetSub(name string) {
	c.Name = name
}

func CallReflect(c interface{}) {
	//rtyp := reflect.TypeOf(c)
	rVal := reflect.ValueOf(c)

	// 获取原始值
	iv := rVal.Interface()
	v ,ok := iv.(*Call)
	if ok {
		fmt.Println("v.Name= ",v.Name)
		fmt.Println("v.Num1= ",v.Num1)
		fmt.Println("v.Num2= ",v.Num2)
	}

	valType := rVal.Kind()
	fmt.Println("valType",valType)
	// 获取字段数
	numField := rVal.Elem().NumField()
	fmt.Println("numFiled= ",numField)

	// 遍历字段
	for i := 0;i<numField;i++{
		filed := rVal.Elem().Field(i)
		fmt.Printf("Filed %d %v\n",i, filed)
	}

	// 获取结构体方法
	//numMoth := rVal.Elem().NumMethod()  // 获取方法数量不需要Elem
	numMethod := rVal.NumMethod()
	fmt.Println("numMethod= ",numMethod)

	// 改变原结构体字段的值
	rVal.Elem().FieldByName("Num1").SetInt(8)
	rVal.Elem().FieldByName("Num2").SetInt(5)

	iMod := rVal.Interface()
	v,ok = iMod.(*Call)
	if ok {
		fmt.Println("Mod: v.Name= ",v.Name)
		fmt.Println("Mod: v.Num1= ",v.Num1)
		fmt.Println("Mod: v.Num2= ",v.Num2)
	}
	// 遍历字段 ergodic
	fmt.Println("遍历字段 ======")
	for i := 0;i<numField;i++{
		filed := rVal.Elem().Field(i)
		fmt.Printf("Filed %d %v\n",i, filed)
	}

	str := fmt.Sprintf("%s 完成了减法运算 %d - %d = %d\n",v.Name, v.Num1,v.Num2,v.Num1-v.Num2)
	fmt.Println(str)


}
