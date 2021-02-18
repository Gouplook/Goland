/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/3 15:03
@Description:

*********************************************/
package digitalsignature

import (
	"fmt"
	"testing"
)

func TestSignatureRSA(t *testing.T) {
	src := []byte("大学之道，在明明德，在亲民，在止於至善")
	signText := SignatureRSA(src,"private.pem")
	fmt.Println("SignText :")
	fmt.Println(signText)
	b := VerifyRSA(src,signText,"public.pem")

	fmt.Printf("%v\n",b)
}