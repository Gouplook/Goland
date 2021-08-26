/**
 * @Author: yinjinlin
 * @File:  gopdf
 * @Description:
 * @Date: 2021/7/19 上午10:50
 */

package gopdf

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)


func loremList() []string {
	return []string{
		"套餐AAAA",
		"20次",
		"单项目AAAA",
		"30次",
		"单项目BBB",
		"40次",
	}
}

type HeTongSingleLists struct {
	SingleName string // 单项目名称
	Num        int    // 单项目次数
}


// 获取切片类型 []string
func HeTongSingle() []string {
	lists := make([]HeTongSingleLists, 3)
	lists[0].SingleName = "木桶足浴"
	lists[0].Num = 22
	lists[1].SingleName = "单BBBBB"
	lists[1].Num = 33
	lists[2].SingleName = "单CCCCCC"
	lists[2].Num = 44

	fmt.Println(lists)

	slic := make([]string, 0)
	var numStr string
	for _, v := range lists {
		numStr = strconv.Itoa(v.Num)
		slic = append(slic, v.SingleName)
		slic = append(slic, numStr)
	}

	return slic
}

func CreateExcel(pdf *gofpdf.Fpdf, getY float64) {

	const (
		colCount = 2    // 表格列数
		colWd    = 60.0 // 单元格行宽
		marginH  = 20.0 // 字符间距
		lineHt   = 5.0  // 行高度
		cellGap  = 2.0  // 单元格的表格表格边距
	)

	type cellType struct {
		str string
		list []string
		ht   float64
	}

	var (
		cellList [colCount]cellType
		cell     cellType
	)

	// 表格标题
	header := [colCount]string{"项目", "次数"}
	alignList := [colCount]string{"C", "C"}
	strList := loremList()
	pdf.SetMargins(marginH, 15, marginH)
	pdf.SetFont("simfang", "", 10)

	pdf.SetXY(20, getY)
	// 设置表格第一行的样式
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 250, 250)  // 背景颜色
	for colJ := 0; colJ < colCount; colJ++ {
		pdf.CellFormat(colWd, 10, header[colJ], "1", 0, "CM", true, 0, "")
	}
	pdf.Ln(-1)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 250, 250)
	y := pdf.GetY()
	count := 0
	// 表格行数
	for rowJ := 0; rowJ < 3; rowJ++ {
		maxHt := lineHt
		// 计算单元格高度
		for colJ := 0; colJ < colCount; colJ++ {
			count++
			if count > len(strList) {
				count = 1
			}
			cell.str = strings.Join(strList[count-1:count], "")
			cell.list = pdf.SplitText(cell.str, colWd-cellGap*2)
			cell.ht = float64(len(cell.list)) * lineHt
			if cell.ht > maxHt {
				maxHt = cell.ht
			}
			cellList[colJ] = cell
		}
		// 循环渲染每个单元格
		x := marginH
		for colJ := 0; colJ < colCount; colJ++ {
			pdf.Rect(x, y, colWd, maxHt+cellGap*2, "D")
			cell = cellList[colJ]
			cellY := y + cellGap + (maxHt-cell.ht)/2
			for splitJ := 0; splitJ < len(cell.list); splitJ++ {
				pdf.SetXY(x+cellGap, cellY) // 60-10
				pdf.CellFormat(colWd-cellGap*2, lineHt, string(cell.list[splitJ]), "", 0, alignList[colJ], false, 0, "")
				cellY += lineHt
			}
			x += colWd
		}
		y += maxHt + cellGap*2

	}

}

func CreatePdf() {
	// 初始化对象（纸张）
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// 获取纸张大小
	w, h := pdf.GetPageSize()
	fmt.Printf("pdf size, w:%.2f, h:%.2f", w, h) // pdf size, w:148.5, h:210.00

	// 添加中文字体
	pdf.AddUTF8Font("simfang", "", "simfang.ttf")
	fornSize := 12.0

	// 标题
	title := "预付卡合同"
	pdf.SetXY(0, 10)
	// 设置字体，字体需要下载
	pdf.SetFont("simfang", "", 20)
	pdf.CellFormat(w, 10, title, "", 1, "TC", false, 0, "")

	wi := pdf.GetStringWidth(title)
	// fmt.Println(wi)

	// 合同编号
	contractNo := "KC2021-0001"
	pdf.SetXY(0, wi/2) // 这种标题居中位置
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, "合同编号："+contractNo, "", 0, "TR", false, 0, "")

	// 设置页脚
	pdf.SetFooterFunc(func() {
		pdf.SetY(-10)
		pdf.CellFormat(0, 10, fmt.Sprintf("当前 第%d页，共{nb}页", pdf.PageNo()), "", 0, "C", false, 0, "")
	})
	pdf.AliasNbPages("")

	// 主体
	chapterBody := func(fileStr string, line float64) {
		textStr, err := ioutil.ReadFile(fileStr)
		if err != nil {
			pdf.SetError(err)
		}

		pdf.SetFont("simfang", "", fornSize)
		pdf.SetXY(20, line)
		str := string(textStr)

		// 替换（采用文信息读取，进行替换）
		out := strings.Replace(str, "UserName", "张三", -1)
		out = strings.Replace(out, "ShopName", "宝山店", -1)
		out = strings.Replace(out, "CardPrice", "500", -1)
		out = strings.Replace(out, "ServicePeriod", "6", -1)
		out = strings.Replace(out, "IsAllSingle", "适用全部", -1)
		// 输出对齐文本（此函数一次性可以插入多行）
		pdf.MultiCell(0, 8, out, "", "", false)
		pdf.Ln(-1)
	}

	// 主体第二段，采用分段进行拼接
	chapterBody2 := func(fileStr string, line float64) {
		textStr, err := ioutil.ReadFile(fileStr)
		if err != nil {
			pdf.SetError(err)
		}

		pdf.SetFont("simfang", "", fornSize)
		pdf.SetXY(20, line)
		str := string(textStr)

		// 替换
		out := strings.Replace(str, "PayRealPrice", "388", -1)
		out = strings.Replace(out, "BuyCardDiscountDesc", "已经优惠", -1)
		out = strings.Replace(out, "UserContactCall", "021-66669990", -1)
		out = strings.Replace(out, "ShopContactCall", "021-77088889", -1)
		out = strings.Replace(out, "PayTimeStr", "2021-07021", -1)
		out = strings.Replace(out, "UserName", "张三", -1)
		out = strings.Replace(out, "ShopName", "宝山店", -1)

		r := []rune(out)
		// 输出对齐文本
		// pdf.MultiCell(0, 8, out, "", "", false)
		pdf.MultiCell(0, 8, string(r), "", "", false)
		pdf.Ln(-1)
	}

	chapterBody("./3.txt", 40)
	// 中间加表格
	CreateExcel(pdf, pdf.GetY()-8)

	chapterBody2("./4.txt", pdf.GetY()+8)

	// 读取文件时，注意文件路径问题
	err := pdf.OutputFileAndClose("../pdf/12.pdf")
	if err != nil {
		panic(err)
	}

}


//
func read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}


// 将文件显示到网页上
func Web() {
	CreatePdf()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/pdf")
		// 支持网页下载功能
		// fileContentDisposition := "attachment;filename=\"" + "00001.pdf" + "\""
		// writer.Header().Set("Content-Disposition",fileContentDisposition)
		content, err := read("002.pdf")
		if err != nil {
			log.Println(err.Error())
		}
		writer.Write(content)
	})
	defer func() {
		del := os.Remove("002.pdf")
		if del != nil {
			fmt.Println(del)
		}
	}()
	log.Fatal(http.ListenAndServe(":8090", nil))

}


func GOPdf(){


}