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
	"strings"
)

func PDF() {

	// p： 纵向
	pdf := gofpdf.New("P", "mm", "A5", "")
	pdf.AddPage()
	// pdf.SetFont("Arial","",56)
	// pdf.Write(10,"Herrllo-----")
	// for fontSize := 4; fontSize < 40; fontSize += 10 {
	// 	pdf.SetFont("Arial", "", 12)
	// 	pdf.SetXY(0, float64(fontSize))
	// 	pdf.Cell(40, 10, "Hello World")

	// }

	// 添加中文字体
	pdf.AddUTF8Font("simfang", "", "simfang.ttf")

	// 标题
	title := "预付卡合同"
	pdf.SetXY(0, 10)
	pdf.SetFont("simfang", "", 20)
	pdf.CellFormat(0, 100, title, "", 0, "TC", false, 0, "")
	// pdf.SetTitle(title,false)
	// pdf.SetAuthor("Jules Verne",false)

	// 合同编号
	contractNo := "KC2021-0001"
	pdf.SetXY(0, 20)
	pdf.SetFont("simfang", "", 10)
	pdf.CellFormat(0, 100, "合同编号："+contractNo, "", 0, "TR", false, 0, "")

	// 正文
	partyA := "张三"
	partyB := "康存集团"
	pdf.SetXY(20, 40)
	pdf.SetFont("simfang", "", 10)
	pdf.CellFormat(0, 100, "甲方："+partyA+"（购卡人也称持卡人）", "", 0, "TL", false, 0, "")

	pdf.SetXY(20, 45)
	pdf.SetFont("simfang", "", 10)
	pdf.CellFormat(0, 100, "乙方："+partyB+"（发售卡人）", "", 0, "TL", false, 0, "")

	// Body 体
	// strBody1 := "甲方购买由乙方发行的单用途商业预付卡(以下简称单用途卡)，并由乙方为持卡人提供"
	// pdf.SetXY(20,50)
	// pdf.SetFont("simfang","",10)
	// pdf.CellFormat(0,100,strBody1,"",0, "TL",false,0,"")
	//
	// strBody2 := "刷卡消费以获取商品或者服务。经甲乙双方友好协商，就有关服务事宜形成的关系约定"
	// pdf.SetXY(20,55)
	// pdf.SetFont("simfang","",10)
	// pdf.CellFormat(0,100,strBody2,"",0, "TL",false,0,"")
	//
	// strBody3 := "如下："
	// pdf.SetXY(20,60)
	// pdf.SetFont("simfang","",10)
	// pdf.CellFormat(0,100,strBody3,"",0, "TL",false,0,"")
	//
	// strBody4 := "第一条单用途卡的功能、金额、注意事项"
	// pdf.SetXY(20,65)
	// pdf.SetFont("simfang","",10)
	// pdf.CellFormat(0,100,strBody4,"",0, "TL",false,0,"")
	//
	// strBody5 := "2.卡值金额："
	//
	// valueStr := "500"
	// pdf.SetXY(20,70)
	// pdf.SetFont("simfang","",10)
	// pdf.CellFormat(0,100,strBody5+valueStr,"",0, "TL",false,0,"")

	// // 下划线
	// pdf.SetXY(0,50)
	// pdf.SetFont("Arial","U",16)
	// pdf.CellFormat(0,100,"card yfraslfdafsf","",0, "TL",false,0,"")

	// // pdf.CellFormat()
	//
	// // 添加第二页
	// pdf.AddPage()
	// pdf.SetFont("Arial","",12)
	// pdf.Write(0,"di er ye....")

	// 设置页脚
	pdf.SetFooterFunc(func() {
		pdf.SetY(-10)
		pdf.CellFormat(0, 10, fmt.Sprintf("当前 第%d页，共{nb}页", pdf.PageNo()), "", 0, "C", false, 0, "")
	})
	pdf.AliasNbPages("")

	err := pdf.OutputFileAndClose("001.pdf")
	if err != nil {
		panic(err)
	}

}

func Boby() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	w, h := pdf.GetPageSize()
	fmt.Printf("pdf size, w:%.2f, h:%.2f", w, h) // pdf size, w:210.00, h:297.00

	pdf.AddUTF8Font("simfang", "", "simfang.ttf")
	titleStr := "合同"
	pdf.SetTitle(titleStr, true)
	// pdf.SetAuthor("世界", true)

	pdf.SetHeaderFuncMode(func() {
		pdf.SetFont("simfang", "", 15)
		wd := pdf.GetStringWidth(titleStr) + 6
		pdf.SetY(0.6)            // 先要设置 Y，然后再设置 X。否则，会导致 X 失效
		pdf.SetX((210 - wd) / 2) // 水平居中的算法

		// pdf.SetDrawColor(0, 80, 180)  //frame color
		// pdf.SetFillColor(230, 230, 0) //background color
		// pdf.SetTextColor(220, 50, 50) //text color
		pdf.SetFillColor(255, 255, 255)
		pdf.SetTextColor(0, 0, 0)

		pdf.SetLineWidth(1)

		pdf.CellFormat(wd, 10, titleStr, "", 1, "CM", true, 0, "")
		// 第 5 个参数，实际效果是：指定下一行的位置

		pdf.Ln(5)

	}, false)

	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.SetTextColor(128, 128, 128)
		pdf.CellFormat(
			0, 5,
			fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "",
		)
	})

	// 标题
	chapterTitle := func(chapNum int, titleStr string) {
		pdf.SetFont("Arial", "", 12)
		pdf.SetFillColor(200, 220, 255) // background color
		pdf.CellFormat(
			0, 6,
			fmt.Sprintf("Chapter %d : %s", chapNum, titleStr),
			"", 1, "L", true, 0, "",
		)
		pdf.Ln(2)

	}
	// 主体
	chapterBody := func(fileStr string) {
		textStr, err := ioutil.ReadFile(fileStr)
		if err != nil {
			pdf.SetError(err)
		}

		pdf.SetFont("simfang", "", 12)

		// 输出对齐文本
		pdf.MultiCell(0, 5, string(textStr), "", "", false)

		pdf.Ln(-1)

		pdf.SetFont("simfang", "", 0)
		pdf.Cell(0, 5, "(end of excerpt)")
	}

	// 印刷每一页
	printChapter := func(chapNum int, titleStr, fileStr string) {
		pdf.AddPage()
		chapterTitle(chapNum, titleStr)

		chapterBody(fileStr)
	}

	printChapter(1, "Once", "./1.txt")
	printChapter(2, "Two", "./2.txt")

	if err := pdf.OutputFileAndClose("1.pdf"); err != nil {
		panic(err.Error())
	}

}

func Contract() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	w, h := pdf.GetPageSize()
	fmt.Printf("pdf size, w:%.2f, h:%.2f", w, h) // pdf size, w:148.5, h:210.00

	// 添加中文字体
	pdf.AddUTF8Font("simfang", "", "simfang.ttf")
	fornSize := 12.0

	// 标题
	title := "预付卡合同"
	pdf.SetXY(0, 10)
	pdf.SetFont("simfang", "", 20)
	pdf.CellFormat(w, 10, title, "", 1, "TC", false, 0, "")

	wi := pdf.GetStringWidth(title)
	fmt.Println(wi)

	// 合同编号
	contractNo := "KC2021-0001"
	pdf.SetXY(0, wi/2)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, "合同编号："+contractNo, "", 0, "TR", false, 0, "")

	// 正文
	partyA := "张三"
	partyB := "康存集团"
	pdf.SetXY(20, 40)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(w, 10, "甲方："+partyA+"（购卡人也称持卡人）", "", 1, "TL", false, 0, "")

	pdf.SetXY(20, 45)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(w, 10, "乙方："+partyB+"（发售卡人）", "", 1, "TL", false, 0, "")

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
		// 输出对齐文本
		pdf.MultiCell(0, 5, string(textStr), "", "", false)
		pdf.Ln(-1)
	}

	// y50 := pdf.GetY()
	// fmt.Println("y50=",y50)

	// chapterBody("./1.txt", 50)
	chapterBody("./1.txt", pdf.GetY()-5)

	strBody5 := "2.卡值金额："
	valueStr := "500"
	// pdf.SetXY(20, 70)

	pdf.SetXY(20, pdf.GetY()-5)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, strBody5+valueStr, "", 0, "TL", false, 0, "")

	strBody6 := "3.有效期限："
	valueDate := "2021-07-21"
	// pdf.SetXY(20, 75)
	y75 := pdf.GetY()
	fmt.Println("y70", y75)
	pdf.SetXY(20, pdf.GetY()+5)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, strBody6+valueDate, "", 0, "TL", false, 0, "")

	strBody7 := "4.适用项目："
	valueStr2 := "适用全部"
	pdf.SetXY(20, pdf.GetY()+5)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, strBody7+valueStr2, "", 0, "TL", false, 0, "")

	// 适用部分项目，需要插入表格

	Excel2(pdf, pdf.GetY()+5)

	strBody8 := "5.购卡金额:  "
	valueAmount := "132.90"
	// pdf.SetXY(20, 85)
	pdf.SetXY(20, pdf.GetY()+10)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, strBody8+valueAmount, "", 0, "TL", false, 0, "")
	//
	strBody9 := "6.可享优惠: "
	valueStr3 := "已经优惠"
	pdf.SetXY(20, pdf.GetY()+5)
	pdf.SetFont("simfang", "", 10)
	pdf.CellFormat(0, 100, strBody9+valueStr3, "", 0, "TL", false, 0, "")

	// chapterBody("./2.txt", 90)
	chapterBody("./2.txt", pdf.GetY()+5)
	// pdf.AddPage()

	pdf.Ln(10)
	// 尾部
	partyA2 := "甲    方："
	partyB2 := "乙    方："
	partyAvule := "张三"

	yy20 := pdf.GetY()
	fmt.Println("yy20", yy20)
	// pdf.SetXY(30, 20)
	pdf.SetXY(30, pdf.GetY())
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, partyA2+partyAvule, "", 0, "TL", false, 0, "")

	// pdf.SetXY(110, 20)
	pdf.SetXY(110, pdf.GetY())
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, partyB2+partyAvule, "", 0, "TL", false, 0, "")

	tel := "联系方式："
	telA := " 13012349999"
	yy30 := pdf.GetY()
	fmt.Println("yy20", yy30)
	// pdf.SetXY(30, 25)
	pdf.SetXY(30, pdf.GetY()+5)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, tel+telA, "", 0, "TL", false, 0, "")

	pdf.SetXY(110, pdf.GetY())
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, tel+telA, "", 0, "TL", false, 0, "")
	//
	signDate := "签约日期："

	// pdf.SetXY(30, 30)
	pdf.SetXY(30, pdf.GetY()+5)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, signDate+valueDate, "", 0, "TL", false, 0, "")

	pdf.SetXY(110, pdf.GetY())
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, signDate+valueDate, "", 0, "TL", false, 0, "")

	err := pdf.OutputFileAndClose("0218.pdf")
	if err != nil {
		panic(err)
	}

	// del := os.Remove("0218.pdf")
	// if del != nil {
	// 	fmt.Println(del)
	// }

}

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

func Excel() {

	const (
		colCount = 2
		colWd    = 60.0
		// marginH  = 15.0
		marginH = 45.0
		lineHt  = 5.5
		cellGap = 2.0
	)

	type cellType struct {
		str  string
		list [][]byte
		ht   float64
	}

	var (
		cellList [colCount]cellType
		cell     cellType
	)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("simfang", "", "simfang.ttf")
	header := [colCount]string{"项目", "次数"}
	alignList := [colCount]string{"L", "C"}
	strList := loremList()
	pdf.SetMargins(marginH, 15, marginH)
	pdf.SetFont("simfang", "", 10)
	pdf.AddPage()

	// 设置表格第一行的样式
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 250, 250)
	for colJ := 0; colJ < colCount; colJ++ {
		pdf.CellFormat(colWd, 10, header[colJ], "1", 0, "CM", true, 0, "")
	}
	pdf.Ln(-1)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 250, 250)
	y := pdf.GetY()
	count := 0
	// 表格行数
	for rowJ := 0; rowJ < 2; rowJ++ {
		maxHt := lineHt
		// 计算单元格高度
		for colJ := 0; colJ < colCount; colJ++ {
			count++
			if count > len(strList) {
				count = 1
			}
			cell.str = strings.Join(strList[colJ:count], " ")
			cell.list = pdf.SplitLines([]byte(cell.str), colWd-cellGap*2) // 60-10
			f := len(cell.list)
			fmt.Println(f)
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
				pdf.SetXY(x+cellGap, cellY)
				pdf.CellFormat(colWd-cellGap*2, lineHt, string(cell.list[splitJ]), "", 0, alignList[colJ], false, 0, "")
				cellY += lineHt
			}
			x += colWd
		}
		y += maxHt + cellGap*2

	}

	if err := pdf.OutputFileAndClose("122.pdf"); err != nil {
		panic(err.Error())
	}

}

func Excel2(pdf *gofpdf.Fpdf, getY float64) {

	const (
		colCount = 2
		colWd    = 60.0
		marginH = 20.0  //
		lineHt  = 5.0
		cellGap = 2.0
	)

	type cellType struct {
		str  string
		//list [][]byte
		list []string
		ht   float64
	}

	var (
		cellList [colCount]cellType
		cell     cellType
	)

	// pdf := gofpdf.New("P", "mm", "A4", "")
	// pdf.AddUTF8Font("simfang", "", "simfang.ttf")
	header := [colCount]string{"项目", "次数"}
	alignList := [colCount]string{"C", "C"}
	strList := loremList()
	pdf.SetMargins(marginH, 15, marginH)
	pdf.SetFont("simfang", "", 10)

	pdf.SetXY(20, getY)
	// 设置表格第一行的样式
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 250, 250)
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
			cell.str = strings.Join(strList[count-1:count], " ")
			//cell.list = pdf.SplitLines([]byte(cell.str), colWd-cellGap*2) // 60-10
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

	// if err := pdf.OutputFileAndClose("122.pdf");err != nil {
	// 	panic(err.Error())
	// }

}

// 思路：
// 前端过来参数参数
// 预览不上传  goland 生成pdf预览

func Replace2() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	w, h := pdf.GetPageSize()
	fmt.Printf("pdf size, w:%.2f, h:%.2f", w, h) // pdf size, w:148.5, h:210.00

	// 添加中文字体
	pdf.AddUTF8Font("simfang", "", "simfang.ttf")
	fornSize := 12.0

	// 标题
	title := "预付卡合同"
	pdf.SetXY(0, 10)
	pdf.SetFont("simfang", "", 20)
	pdf.CellFormat(w, 10, title, "", 1, "TC", false, 0, "")

	wi := pdf.GetStringWidth(title)
	fmt.Println(wi)

	// 合同编号
	contractNo := "KC2021-0001"
	pdf.SetXY(0, wi/2)
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

		// 替换
		out := strings.Replace(str, "UserName", "张三", -1)
		out = strings.Replace(out, "ShopName", "宝山店", -1)
		out = strings.Replace(out, "CardPrice", "500", -1)
		out = strings.Replace(out, "ServicePeriod", "6", -1)
		out = strings.Replace(out, "IsAllSingle", "适用全部", -1)
		// 输出对齐文本
		pdf.MultiCell(0, 8, out, "", "", false)
		pdf.Ln(-1)
	}

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

		// // 输出对齐文本
		pdf.MultiCell(0, 8, out, "", "", false)
		pdf.Ln(-1)
	}


	chapterBody("./3.txt", 40)
	// 中间加表格
	Excel2(pdf,pdf.GetY()-8)

	chapterBody2("./4.txt",pdf.GetY()+10)



	err := pdf.OutputFileAndClose("0021.pdf")
	if err != nil {
		panic(err)
	}

}
