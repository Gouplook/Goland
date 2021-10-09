/**
 * @Author: yinjinlin
 * @File:  gotoexcel.go
 * @Description: 查找到数据，生产excel表，导出
 * @Date: 2021/10/9 上午10:56
 */

package goexcel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// 思路：
// 调用接口，传过来数据生产excel表格，根据选择的页面生产excel数据


func CreatExcel(){
	// categories := map[string]string{
	// 	"A2": "Small", "A3": "Normal", "A4": "Large",
	// 	"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	// values := map[string]int{
	// 	"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}

	excelTite := map[string]string {
		"A1": "序号",
		"B1": "订单编号",
		"C1": "门店名称",
		"D1": "订单类型",
		"E1": "城市",
		"F1": "总金额",
		"G1": "支付渠道",
		"H1": "支付方式",
		"I1": "付款时间",
		"J1": "支付手续费",
		"K1": "存管比例",
		"L1": "存管金额",
		"M1": "保费",
		"N1": "续保保费",
		"O1": "平台手续费",
		"P1": "商家应得金额",

	}

	expMap := []map[string]string{

	}
	// 新建工作薄
	f := excelize.NewFile()
	// f.SetColWidth("Sheet1","A","H",20)
	// f.SetRowHeight("sheet1",1,10)
	// for k, v := range categories {
	// 	f.SetCellValue("Sheet1", k, v)
	// }
	// for k, v := range values {
	// 	f.SetCellValue("Sheet1", k, v)
	// }

	// 设置单元格样式
	// style,err := f.NewStyle()
	// f.SetCellStyle()

	for k, v := range excelTite {
		f.SetCellValue("Sheet1",k, v)
	}



	// []map[sting]string{}
	length := len(expMap)
	for i := 0;i<length;i++{
		for k, v := range expMap[i] {
			f.SetCellValue("Sheet1",k, v)
		}
	}


	st,_ := f.CalcCellValue("sheet","N1")
	fmt.Println(st)

	excelName := "Book2" + ".xlsx"

	// Save spreadsheet by the given path.
	if err := f.SaveAs(excelName); err != nil {
		fmt.Println(err)
	}

	// 文件先生存，上传后就直接删除。

}


func CreteExcel2(){

}

