/**
 * @Author: yinjinlin
 * @File:  goexcel
 * @Description:
 * @Date: 2021/6/22 下午5:19
 */

package dbtool

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

type Kc_cnaps struct {
	Id       int
	BankCode string
	BankName string
	// Ys   string
	// Gg   string
	// Ddh  string
	// Xfz  string
	// Lxfs string
	// Jhrq string
	// Shrq string
}

func EXcelInit() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/jkd_cards?charset=utf8", 30)
	orm.RegisterModel(new(Kc_cnaps))
	_ = orm.RunSyncdb("default", false, true)
}

func ExcelMain() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: xlsx pathname sheetname")
		os.Exit(1)
	}

	xlsxFile, err := xlsx.OpenFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sheet := xlsxFile.Sheet[os.Args[2]]
	if sheet == nil {
		fmt.Println("表单名不存在")
		os.Exit(1)

	}

	cols := len(sheet.Cols)
	sleet := make([][]string, cols)
	oksleet := make([][]string, len(sheet.Rows)-1)
	log.Println("列数：", cols, "   行数：", len(sheet.Rows))

	var line int
	//rows := len(sheet.Rows)
	//for i := 0; i < cols; i++ {
	// sleet[i] = append(sleet[i], []string)
	//}

	// lname := []string{"客户名称", "品名", "颜色", "规格", "订单号", "消费者", "联系方式", "交货日期", "审核日期"}
	lname := []string{"BANK_CODE", "LNAME"}

	for _, title := range lname {
		//title := getStdinInput("请输入列名：")
		//if title == "" {
		// fmt.Println(title)
		// continue
		//}

		titleColIndex := findColByTitle(sheet, title)
		if titleColIndex == -1 {
			fmt.Println("列名不存在")
			continue
		}

		rowLen := len(sheet.Rows)
		result := []string{}
		for rowIndex := 1; rowIndex < rowLen; rowIndex++ {
			content := sheet.Cell(rowIndex, titleColIndex).String()
			result = append(result, content)
			sleet[line] = append(sleet[line], content)

			//fmt.Println(sleet)
			//if line == 8 {
			// fmt.Println(sleet)
			//}

		}

		log.Println("line:", line)

		if line == cols-1 {
			for i := 0; i < len(sheet.Rows)-1; i++ {
				for j := 0; j < cols; j++ {
					oksleet[i] = append(oksleet[i], sleet[j][i])
					//log.Println(oksleet[i])
				}
			}
			//fmt.Println(oksleet)
			for _, v := range oksleet {
				msg := Kc_cnaps{}
				// msg.Id = v[0]
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
