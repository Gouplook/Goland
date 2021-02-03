/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/2 15:11
@Description:

*********************************************/
package asymmetricEncrypt

import (
	"crypto/md5"
	"fmt"
	"testing"
)



// 生产密钥
func TestGetPrivate(t *testing.T){
	GetRsakey(2048)
}


//
func TestRSAEncrypt(t *testing.T) {
	src := []byte("EncryptPEMBlock使用指定的密码3333 打开文件")

	fmt.Println("加密前==： ")
	cipherText := RSAEncrypt(src,"public.pem")

	plainText := RSADencrpt(cipherText,"private.pem")

	fmt.Println("解密后==： ")
	fmt.Println(string(plainText))

	data := []byte("These pretzels are making me thirsty.")
	fmt.Printf("%v",md5.Sum(data))
}