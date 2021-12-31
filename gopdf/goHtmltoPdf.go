/**
 * @Author: yinjinlin
 * @File:  goHtmltoPdf
 * @Description:
 * @Date: 2021/12/28 下午1:18
 */

package gopdf

import (
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
)

//  HTML 转化pdf
//  jdx006884011015-1-1

// 开发摘要：
// 先从master 切一个分支出来开发，最后合并到master和dev上 删除分支
// 涉及到的容器有：rpcCard、rpcOder、 apiCrard.
// 发卡业务、卡包业务。
// 注意问题，开发2期，是否兼顾3期。还是先开发2期，后面在开发三期。


const (
	OrientationLandscape = "Landscape" // Landscape mode
	OrientationPortrait  = "Portrait"  // Portrait mode
)

// 涉及问题点
// 分页问题。

func HtmlToPdf() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		log.Fatal(err)
	}

	//
	// `card_range` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '卡项使用范围,1=仅限指定门店使用,2=直营连锁门店通用,3=加盟门店通用,4=其它合作门店通用',
	// 	`card_limits` varbinary(20) DEFAULT NULL COMMENT ' 多个用逗号隔开，卡项使用权限 1=本人使用 2=家人使用 3= 朋友使用 4= 同事使用，5=任何人使用 6=可转让 7=可共享',
	// 	`special_agreements` varbinary(50) DEFAULT NULL COMMENT '特别约定',
	// 	PRIMARY KEY (`id`),


	// F_card_range         string `default:"card_range"`
	// F_card_limits        string `default:"card_limits"`
	// F_special_agreements string `default:"special_agreements"`


	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA5)
	pdfg.Orientation.Set(OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage("./html.html")



	// Set options for this page
	// page.FooterRight.Set("[page]")
	str := "合同编号：HT000001"
	// page.HeaderCenter.Set(str)
	page.HeaderLeft.Set(str)
	page.HeaderLine.Set(true)
	page.HeaderFontSize.Set(7)
	page.HeaderSpacing.Set(3)

	page.FooterCenter.Set("第 [page]页")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.9)


	// Add to document  // 4
	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err  = pdfg.Create()
	if err != nil {

	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./007003.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done

}
