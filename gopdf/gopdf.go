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

func PDF(){

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
	// pdf.SetXY(0,10)
	// pdf.SetFont("simfang","B",50)
	// pdf.CellFormat(0,100,title,"",0, "TC",false,0,"")
	pdf.SetTitle(title,false)
	pdf.SetAuthor("Jules Verne",false)

	// 合同编号
	contractNo := "KC2021-0001"
	pdf.SetXY(0,20)
	pdf.SetFont("simfang","",10)
	pdf.CellFormat(0,100,"合同编号："+contractNo,"",0, "TR",false,0,"")

	// 正文
	partyA :="张三"
	partyB := "康存集团"
	pdf.SetXY(20,40)
	pdf.SetFont("simfang","",10)
	pdf.CellFormat(0,100,"甲方：" + partyA + "（购卡人也称持卡人）","",0, "TL",false,0,"")

	pdf.SetXY(20,50)
	pdf.SetFont("simfang","",10)
	pdf.CellFormat(0,100,"乙方：" + partyB + "（发售卡人）","",0, "TL",false,0,"")



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
		pdf.CellFormat(0,10,fmt.Sprintf("当前 第%d页，共 {nb} 页", pdf.PageNo()),"",0,"C",false,0,"")
	})
	pdf.AliasNbPages("")

	err := pdf.OutputFileAndClose("001.pdf")
	if err != nil {
		panic(err)
	}

}


func Boby(){
	pdf := gofpdf.New("P", "mm", "A4", "")
	w, h := pdf.GetPageSize()
	fmt.Printf("pdf size, w:%.2f, h:%.2f", w, h) //pdf size, w:210.00, h:297.00

	pdf.AddUTF8Font("simfang", "", "simfang.ttf")
	titleStr := "预付卡合同"
	pdf.SetTitle(titleStr, false)
	pdf.SetAuthor("Jules Verne", false)


	pdf.SetHeaderFuncMode(func() {
		pdf.SetFont("simfang", "B", 15)
		wd := pdf.GetStringWidth(titleStr) + 6
		pdf.SetY(0.6)            //先要设置 Y，然后再设置 X。否则，会导致 X 失效
		pdf.SetX((210 - wd) / 2) //水平居中的算法

		pdf.SetDrawColor(0, 80, 180)  //frame color
		pdf.SetFillColor(230, 230, 0) //background color
		pdf.SetTextColor(220, 50, 50) //text color

		pdf.SetLineWidth(1)

		pdf.CellFormat(wd, 10, titleStr, "1", 1, "CM", true, 0, "")
		//第 5 个参数，实际效果是：指定下一行的位置

		pdf.Ln(5)


	},false)

	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("simfang", "I", 8)
		pdf.SetTextColor(128, 128, 128)
		pdf.CellFormat(
			0, 5,
			fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "",
		)
	})

	//标题
	chapterTitle := func(chapNum int, titleStr string) {
		pdf.SetFont("simfang", "", 12)
		pdf.SetFillColor(200, 220, 255) // background color
		pdf.CellFormat(
			0, 6,
			fmt.Sprintf("Chapter %d : %s", chapNum, titleStr),
			"", 1, "L", true, 0, "",
		)
		pdf.Ln(2)

	}
	//主体
	chapterBody := func(fileStr string) {
		textStr, err := ioutil.ReadFile(fileStr)
		if err != nil {
			pdf.SetError(err)
		}

		pdf.SetFont("simfang", "", 12)

		//输出对齐文本
		pdf.MultiCell(0, 5, string(textStr), "", "", false)

		pdf.Ln(-1)

		pdf.SetFont("simfang", "I", 0)
		pdf.Cell(0, 5, "(end of excerpt)")
	}

	//印刷每一页
	printChapter := func(chapNum int, titleStr, fileStr string) {
		pdf.AddPage()
		chapterTitle(chapNum, titleStr)

		chapterBody(fileStr)
	}

	printChapter(1, "A RUNAWAY REEF", "./1.txt")
	printChapter(2, "THE PROS AND CONS", "./2.txt")


	if err := pdf.OutputFileAndClose("1.pdf"); err != nil {
		panic(err.Error())
	}


}



