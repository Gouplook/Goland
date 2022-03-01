/**
 * @Author: yinjinlin
 * @File:  goexcel
 * @Description:  Excel表格数据读入到mysql数据库中
 * @Date: 2021/6/29 上午10:13
 */

package goexcel

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

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



func  GoExcel() {

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

	// 表格列名称
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
				msg.BankName = v[2]
				msg.BankName = v[3]
				msg.BankName = v[4]

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
		//
		if col.String() == title {
			return titleIndex
		}
	}
	// 5 + 2 + 1
	//

	return -1
}