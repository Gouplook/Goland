package main


import (
	"fmt"
)

// 加法类
type Add struct {
	Object
}

func (a *Add) GetResult() int { // 方法的实现要和接口中方法的声明保持一致
	return a.numA + a.numB
}
func (a *Add) SetData(data ...interface{}) bool {
	// 1: 对数据的个数进行校验。
	var b bool = true
	if len(data) > 2 {
		fmt.Println("参数个数错误！！")
		b = false
	}
	value, ok := data[0].(int)
	if !ok {
		fmt.Println("第一个数类型错误")
		b = false
	}
	value1, ok1 := data[1].(int)
	if !ok1 {
		fmt.Println("第二个数据类型错误")
		b = false
	}
	a.numA = value
	a.numB = value1
	// 2: 类型进行校验。
	return b
}

// 减法类
type Sub struct {
	Object
}

func (s *Sub) SetData(data ...interface{}) bool {
	// 1: 对数据的个数进行校验。
	var b bool = true
	if len(data) > 2 {
		fmt.Println("参数个数错误！！")
		b = false
	}
	value, ok := data[0].(int)
	if !ok {
		fmt.Println("第一个数类型错误")
		b = false
	}
	value1, ok1 := data[1].(int)
	if !ok1 {
		fmt.Println("第二个数据类型错误")
		b = false
	}
	s.numA = value
	s.numB = value1
	// 2: 类型进行校验。
	return b
}
func (s *Sub) GetResult() int {
	return s.numA - s.numB
}

type Object struct {
	numA int
	numB int
}
type Resulter interface {
	GetResult() int
	SetData(data ...interface{}) bool // 完成参数运算的数据的类型校验。
}

// 对象问题
// 1: 定义一个新的类
type OperatorFactory struct {
}

// 2: 创建一个方法，在该方法中完成对象的创建
func (o *OperatorFactory) CreateOperator(op string) Resulter {
	switch op {
	case "+":
		add := new(Add)
		return add
	case "-":
		sub := new(Sub)
		return sub
	default:
		return nil
	}
}
func OperatorWho(h Resulter) int {
	return h.GetResult()
}
func Addupper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}



func main() {

	fmt.Println("78880====")
	//fmt.Println(utils.Age)
	//f := Addupper()
	//fmt.Println(f(1))
	//fmt.Println(f(2))

	//realPrice := 770.0  // 实际支付
	//price := 1000.0 // 面值
	//discout, _ := decimal.NewFromFloat(realPrice).Div(decimal.NewFromFloat(price)).Truncate(2).Float64()
	//fmt.Println(discout)
	//dis := []float64{2.5,2.9,2.5}
	//
	//dis2 := dis
	//sort.Float64s(dis)
	//fmt.Println(dis)
	//min,max := dis[0],dis[len(dis)-1]
	//fmt.Println(min)
	//fmt.Println(max)
	//fmt.Println(dis2)

	// routers（注册service中的路由） -- logic（实现业务逻辑） -- model（数据模型） --- service(对应接口的）
	fmt.Println(3000+2415+2655+3720+6825+7514+8820)
	fmt.Println(869*17)
	fmt.Println(900*17)

}



