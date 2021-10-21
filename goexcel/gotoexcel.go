/**
 * @Author: yinjinlin
 * @File:  gotoexcel.go
 * @Description: 查找到数据，生产excel表，导出
 * @Date: 2021/10/9 上午10:56
 */

package goexcel

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
	"log"
	"net/http"
	"os"
	"strconv"
)

// 思路：
// 调用接口，传过来数据生产excel表格，根据选择的页面生产excel数据

func CreatExcel() {
	// categories := map[string]string{
	// 	"A2": "Small", "A3": "Normal", "A4": "Large",
	// 	"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	// values := map[string]int{
	// 	"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}

	excelTite := map[string]string{
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
		f.SetCellValue("Sheet1", k, v)
	}

	// []map[sting]string{}
	length := len(expMap)
	for i := 0; i < length; i++ {
		for k, v := range expMap[i] {
			f.SetCellValue("Sheet1", k, v)
		}
	}

	st, _ := f.CalcCellValue("sheet", "N1")
	fmt.Println(st)

	excelName := "Book2" + ".xlsx"

	// Save spreadsheet by the given path.
	if err := f.SaveAs(excelName); err != nil {
		fmt.Println(err)
	}

	// 文件先生存，上传后就直接删除。

}

func CreteExcel2() {

	// 新建工作薄
	f := excelize.NewFile()
	// title := map[string]string {
	// 	// "A1": "序号",
	// 	"B1": "订单编号",
	// 	"C1": "门店名称",
	// 	"D1": "订单类型",
	// 	"E1": "城市",
	// 	"F1": "总金额",
	// 	"G1": "支付渠道",
	// 	"H1": "支付方式",
	// 	"I1": "付款时间",
	// 	"J1": "支付手续费",
	// 	"K1": "存管比例",
	// 	"L1": "存管金额",
	// 	"M1": "保费",
	// 	"N1": "续保保费",
	// 	"O1": "平台手续费",
	// 	"P1": "商家应得金额",
	//
	// }

	// title2 := []string{
	// 	"订单编号",
	// 	"门店名称",
	// }

	// for k,v := range title{
	// 	f.SetCellValue("Sheet1",k, v)
	// }

	excelName := "Book5" + ".xlsx"

	// Save spreadsheet by the given path.
	if err := f.SaveAs(excelName); err != nil {
		fmt.Println(err)
	}

}

func Createexcel3() {
	file := xlsx.NewFile()
	xlsx.SetDefaultFont(16,"FangSong")


	sheet, _ := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	// row.SetHeightCM(1) //设置每行的高度
	cell := row.AddCell()
	cell.Value = "haha"
	cell = row.AddCell()
	cell.Value = "xixi"

	err := file.Save("file3.xlsx")
	if err != nil {
		panic(err)
	}
}

func Down(w http.ResponseWriter, r *http.Request) {
	f := excelize.NewFile()
	title := map[string]string{
		"A1": "订单编号",
		"B1": "门店名称",
		"C1": "订单类型",
		"D1": "城市",
		"E1": "总金额",
		"F1": "支付渠道",
		"G1": "支付方式",
		"H1": "付款时间",
		"I1": "支付手续费",
		"J1": "存管比例",
		"K1": "存管金额",
		"L1": "保费",
		"M1": "续保保费",
		"N1": "平台手续费",
		"O1": "商家应得金额",
	}

	for k, v := range title {
		f.SetCellValue("Sheet1", k, v)
	}

	if err := f.SaveAs("Books1224.xlsx"); err != nil {
		fmt.Println(err)
	}
	//
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+"Books1.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	_ = f.Write(w)
}

func CreateExecel4() {
	http.HandleFunc("/", Down) //   设置访问路由
	log.Fatal(http.ListenAndServe(":10090", nil))

}

func getShuju() (shuju Shuju) {
	shuju.Lists = make([]ShujuBase, 0)
	shuju.Lists = []ShujuBase{
		{
			OrderSn:   "S1245",
			ShopId:    1,
			ShopName:  "店面1",
			CityId:    1,
			CityName:  "北京",
			PayChance: 1,
		},
		{
			OrderSn:   "S1243",
			ShopId:    1,
			ShopName:  "店面2",
			CityId:    1,
			CityName:  "上海",
			PayChance: 2,
		},
		{
			OrderSn:   "S1244",
			ShopId:    1,
			ShopName:  "店面3",
			CityId:    1,
			CityName:  "广州",
			PayChance: 1,
		},
	}

	return
}

type ShujuBase struct {
	OrderSn   string //
	ShopId    int
	ShopName  string //
	CityId    int
	CityName  string //
	PayChance int
}

type Shuju struct {
	Lists []ShujuBase
}

func Createexcel5() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "订单编号"
	cell = row.AddCell()
	cell.Value = "门店名称"
	cell = row.AddCell()
	cell.Value = "城市"
	cell = row.AddCell()
	cell.Value = "支付渠道"

	shujuList := getShuju().Lists
	payChance := "支付宝直连"

	for _, v := range shujuList {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = v.OrderSn
		cell = row.AddCell()
		cell.Value = v.ShopName
		cell = row.AddCell()
		cell.Value = v.CityName
		if v.PayChance == 1 {
			payChance = "微信直连"
		} else {
			payChance = "支付宝直连"
		}
		cell = row.AddCell()
		cell.Value = payChance
	}

	if err := file.Save("demo.xlsx"); err != nil {
		fmt.Println(err.Error())
		return
	}

}
type  ExecelData  struct{
	OrderSn   string //
	ShopId    int
	ShopName  string //
	CityId    int
	CityName  string //
	PayChance int
}


func getExcelData(){

}

// func getExcelData()[]ExecelData{
//
// }

type XlsxRow struct {
	Row *xlsx.Row
	Data []string
}

func newRow(row *xlsx.Row,data []string) *XlsxRow {
	return &XlsxRow{
		Row:row,
		Data:data,
	}
}

func (row *XlsxRow) SetRowTitle() error {
	return generateRow(row.Row,row.Data)
}
func (row *XlsxRow) GenerateRow() error {
	return generateRow(row.Row,row.Data)
}

func generateRow(row *xlsx.Row,rowStr []string) error {
	if rowStr == nil {
		return errors.New("no data to generate xlsx!")
	}
	for _,v := range rowStr {
		cell := row.AddCell()
		cell.SetString(v)
	}
	return nil
}

func download(w http.ResponseWriter, file *xlsx.File){
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+"Books1.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	_ = file.Write(w)
}

func CreateExcel(page int, datas []ExecelData) {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	title := make([]string,0)
	title = append(title,"订单编号")
	title = append(title,"门店名称")
	title = append(title,"订单类型")
	title = append(title,"城市")
	title = append(title,"总金额")
	title = append(title,"支付渠道")
	title = append(title,"支付方式")
	title = append(title,"支付手续费")
	title = append(title,"存管比例")
	title = append(title,"存管金额")
	title = append(title,"保费")
	title = append(title,"续保保费")
	title = append(title,"平台手续费")
	title = append(title,"商家应得金额")

	titleRow := sheet.AddRow()

	xlsRow := newRow(titleRow,title)
	err = xlsRow.SetRowTitle()
	if err != nil {
		fmt.Println(err)
		return
	}


	var payChance string
	for _, v := range datas {
		titleRow = sheet.AddRow()
		cell := titleRow.AddCell()
		cell.Value = v.OrderSn
		cell = titleRow.AddCell()
		cell.Value = v.ShopName
		cell = titleRow.AddCell()
		cell.Value = v.CityName
		if v.PayChance == 1 {
			payChance = "微信直连"
		} else {
			payChance = "支付宝直连"
		}
		cell = titleRow.AddCell()
		cell.Value = payChance
	}

	pageStr := strconv.Itoa(page)
	name := "Books" + "_" + pageStr + ".xlsx"
	err = file.Save(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 导出数据


	// 下载完成以后 删除
	defer func() {
		os.Remove(name)
	}()


}

