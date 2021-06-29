package main

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

// import (
// 	"fmt"
// )
//
// // 加法类
// type Add struct {
// 	Object
// }
//
// func (a *Add) GetResult() int { // 方法的实现要和接口中方法的声明保持一致
// 	return a.numA + a.numB
// }
// func (a *Add) SetData(data ...interface{}) bool {
// 	// 1: 对数据的个数进行校验。
// 	var b bool = true
// 	if len(data) > 2 {
// 		fmt.Println("参数个数错误！！")
// 		b = false
// 	}
// 	value, ok := data[0].(int)
// 	if !ok {
// 		fmt.Println("第一个数类型错误")
// 		b = false
// 	}
// 	value1, ok1 := data[1].(int)
// 	if !ok1 {
// 		fmt.Println("第二个数据类型错误")
// 		b = false
// 	}
// 	a.numA = value
// 	a.numB = value1
// 	// 2: 类型进行校验。
// 	return b
// }
//
// // 减法类
// type Sub struct {
// 	Object
// }
//
// func (s *Sub) SetData(data ...interface{}) bool {
// 	// 1: 对数据的个数进行校验。
// 	var b bool = true
// 	if len(data) > 2 {
// 		fmt.Println("参数个数错误！！")
// 		b = false
// 	}
// 	value, ok := data[0].(int)
// 	if !ok {
// 		fmt.Println("第一个数类型错误")
// 		b = false
// 	}
// 	value1, ok1 := data[1].(int)
// 	if !ok1 {
// 		fmt.Println("第二个数据类型错误")
// 		b = false
// 	}
// 	s.numA = value
// 	s.numB = value1
// 	// 2: 类型进行校验。
// 	return b
// }
// func (s *Sub) GetResult() int {
// 	return s.numA - s.numB
// }
//
// type Object struct {
// 	numA int
// 	numB int
// }
// type Resulter interface {
// 	GetResult() int
// 	SetData(data ...interface{}) bool // 完成参数运算的数据的类型校验。
// }
//
// // 对象问题
// // 1: 定义一个新的类
// type OperatorFactory struct {
// }
//
// // 2: 创建一个方法，在该方法中完成对象的创建
// func (o *OperatorFactory) CreateOperator(op string) Resulter {
// 	switch op {
// 	case "+":
// 		add := new(Add)
// 		return add
// 	case "-":
// 		sub := new(Sub)
// 		return sub
// 	default:
// 		return nil
// 	}
// }
// func OperatorWho(h Resulter) int {
// 	return h.GetResult()
// }
// func Addupper() func(int) int {
// 	var n int = 10
// 	return func(x int) int {
// 		n = n + x
// 		return n
// 	}
// }
//
//
//
// func main() {
//
// 	fmt.Println("78880====")
// 	//fmt.Println(utils.Age)
// 	//f := Addupper()
// 	//fmt.Println(f(1))
// 	//fmt.Println(f(2))
//
// 	//realPrice := 770.0  // 实际支付
// 	//price := 1000.0 // 面值
// 	//discout, _ := decimal.NewFromFloat(realPrice).Div(decimal.NewFromFloat(price)).Truncate(2).Float64()
// 	//fmt.Println(discout)
// 	//dis := []float64{2.5,2.9,2.5}
// 	//
// 	//dis2 := dis
// 	//sort.Float64s(dis)
// 	//fmt.Println(dis)
// 	//min,max := dis[0],dis[len(dis)-1]
// 	//fmt.Println(min)
// 	//fmt.Println(max)
// 	//fmt.Println(dis2)
//
// 	// routers（注册service中的路由） -- logic（实现业务逻辑） -- model（数据模型） --- service(对应接口的）
//
// }



type Msg struct {
	Id   int
	BankCode string
	BankName   string
	// Ys   string
	// Gg   string
	// Ddh  string
	// Xfz  string
	// Lxfs string
	// Jhrq string
	// Shrq string
}

func Init() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "root:shutongadmin2459@tcp(192.168.1.246:3306)/jkd_public?charset=utf8", 30)
	orm.RegisterModel(new(Msg))
	err := orm.RunSyncdb("default", false, true)
	log.Println(err)
}

func  main() {

	Init()
	xlsxFile, err := xlsx.OpenFile("./excel1.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sheet := xlsxFile.Sheet["cbd"]
	if sheet == nil {
		fmt.Println("表单名不存在")
		os.Exit(1)
	}

	cols := len(sheet.Cols)
	sleet := make([][]string, cols)
	oksleet := make([][]string, len(sheet.Rows)-1)
	log.Println("列数：", cols, "   行数：", len(sheet.Rows))

	var line int

	lname := []string{"BANK_CODE","LNAME"}

	for _, title := range lname {
		titleColIndex := findColByTitle(sheet, title)
		if titleColIndex == -1 {
			fmt.Println("列名不存在")
			continue
		}

		rowLen := len(sheet.Rows)
		result := []string{}
		for rowIndex := 1; rowIndex < rowLen; rowIndex++ {
			content := (sheet.Cell(rowIndex, titleColIndex).String())
			result = append(result, content)
			sleet[line] = append(sleet[line], content)
		}

		log.Println("line:", line)

		if line == cols-1 {
			for i := 0; i < len(sheet.Rows)-1; i++ {
				for j := 0; j < cols; j++ {
					oksleet[i] = append(oksleet[i], sleet[j][i])
				}
			}
			for _, v := range oksleet {
				msg := Msg{}
				msg.BankCode = v[0]
				msg.BankName = v[1]

				//插入数据表
				o := orm.NewOrm()
				id, err := o.Insert(&msg)
				if err == nil {
					fmt.Println("写入成功,ID为：", id)
				}
			}
		}
		line++
	}
}

func getStdinInput(hint string) string {
	fmt.Print(hint)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func findColByTitle(sheet *xlsx.Sheet, title string) int {
	titleRow := sheet.Rows[0]
	for titleIndex, col := range titleRow.Cells {
		if col.String() == title {
			return titleIndex
		}
	}

	return -1
}