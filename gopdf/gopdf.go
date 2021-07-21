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


	chapterBody("./1.txt", 50)

	strBody5 := "2.卡值金额："
	valueStr := "500"
	pdf.SetXY(20, 70)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, strBody5+valueStr, "", 0, "TL", false, 0, "")

	strBody6 := "3.有效期限："
	valueDate := "2021-07-21"
	pdf.SetXY(20, 75)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, strBody6+valueDate, "", 0, "TL", false, 0, "")

	strBody7 := "4.适用项目："
	valueStr2 := "适用全部"
	pdf.SetXY(20, 80)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, strBody7+valueStr2, "", 0, "TL", false, 0, "")

	strBody8 := "5.购卡金额:  "
	valueAmount := "132.90"
	pdf.SetXY(20, 85)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, strBody8+valueAmount, "", 0, "TL", false, 0, "")
	//
	strBody9 := "6.可享优惠: "
	valueStr3 := "已经优惠"
	pdf.SetXY(20, 90)
	pdf.SetFont("simfang", "", 10)
	pdf.CellFormat(0, 100, strBody9+valueStr3, "", 0, "TL", false, 0, "")


	// 插入表格


	chapterBody("./2.txt",95)
	// pdf.AddPage()

	pdf.Ln(10)
	// 尾部
	partyA2 := "甲    方："
	partyB2 := "乙    方："
	partyAvule := "张三"
	pdf.SetXY(30,20)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, partyA2+partyAvule, "", 0, "TL", false, 0, "")

	pdf.SetXY(110, 20)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, partyB2+partyAvule, "", 0, "TL", false, 0, "")

	tel := "联系方式："
	telA := " 13012349999"
	pdf.SetXY(30, 25)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, tel+telA, "", 0, "TL", false, 0, "")

	pdf.SetXY(110, 25)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, tel+telA, "", 0, "TL", false, 0, "")

	signDate := "签约日期："

	pdf.SetXY(30, 30)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, signDate+valueDate, "", 0, "TL", false, 0, "")

	pdf.SetXY(110, 30)
	pdf.SetFont("simfang", "", fornSize)
	pdf.CellFormat(0, 100, signDate+valueDate, "", 0, "TL", false, 0, "")

	err := pdf.OutputFileAndClose("0216.pdf")
	if err != nil {
		panic(err)
	}

}
