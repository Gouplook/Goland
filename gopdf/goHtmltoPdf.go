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

// HTML 转化pdf
//  jdx006884011015-1-1

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

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage("./007.html")

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)


	// Add to document  // 4
	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err  = pdfg.Create()
	if err != nil {

	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./007.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done

}
