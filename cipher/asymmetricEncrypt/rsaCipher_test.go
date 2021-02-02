/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/2 15:11
@Description:

*********************************************/
package asymmetricEncrypt

import (
	"fmt"
	"testing"
)



// 生产密钥
func TestGetPrivate(t *testing.T){
	GetRsakey(1024)
}


//
func TestRSAEncrypt(t *testing.T) {
	src := []byte("EncryptPEMBlock使用指定的密码、加密算法加密data，返回一个具有指定块类型，保管加密后数据的PEM块")

	fmt.Println("加密前==： ")
	cipherText := RSAEncrypt(src,"public.pem")

	plainText := RSADencrpt(cipherText,"private.pem")

	fmt.Println("解密后==： ")
	fmt.Println(string(plainText) +"88")
}