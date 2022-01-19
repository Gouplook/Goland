/**
 * @Author: yinjinlin
 * @File:  goHtmltoPdf_test
 * @Description:
 * @Date: 2021/12/28 下午2:30
 */

package gopdf

import (
	"fmt"
	"testing"
)

func TestHtmlToPdf(t *testing.T) {
	HtmlToPdf()
}

func TestHtmlTo(t *testing.T) {
	var i = 781
	i = 012345
	newStr := fmt.Sprintf("%011d",i)
	fmt.Println(newStr)
}