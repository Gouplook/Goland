/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/3 17:17
@Description:

*********************************************/
package digitalsignature

import (
	"fmt"
	"testing"
)

func TestEccVerify(t *testing.T) {
	GenerateEcckey()
	src := []byte("格物、致知、诚意、正心、修身、齐家、治国、平天下")
	rText, sText := EccSignature(src,"esprivate.pem")
	bl := EccVerify(src, rText, sText, "espublic.pem")
	fmt.Println(bl)
}