/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/1 15:07
@Description:

*********************************************/
package sysmmertic

import (
	"fmt"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	plainText := "54创建一个使用cbc分组接口AES11"

	key := "abcd1234abcd1234abcd1234"
	// 加密
	fmt.Printf("加密前：%v", plainText )
	fmt.Println()
	cipherByte := AesEncrypt([]byte(plainText),[]byte(key))
	fmt.Println("====", string(cipherByte))


	// 解密
	decText := AesDencrypt(cipherByte,[]byte(key))
	fmt.Println("解密====", string(decText))

}
// OFB 测试
func TestAesDencryptOfb(t *testing.T) {
	plainText := "#OFb54创建一个使用cbc分组接口AES11"
	key := "abcd1234abcd1234abcd1234"
	// 加密
	fmt.Printf("加密前：%v", plainText )
	fmt.Println()
	cipherByte := AesEncryptOfb([]byte(plainText),[]byte(key))
	fmt.Println("====", string(cipherByte))

	// 解密
	//decText := AesDencryptOfb(cipherByte, []byte(key))
	//fmt.Println("解密====", string(decText))

	decText := AesEncryptOfb(cipherByte, []byte(key))
	fmt.Println("解密====", string(decText))


}

// CTR测试
func TestAesDencryptCTR(t *testing.T) {
	plainText :=[]byte("使用CTR分组接口")
	key := []byte("abcd1234abcd1234abcd1234")
	fmt.Printf("加密前：%v", string(plainText) )
	fmt.Println()
	cipherByte := AesEncryptCTR(plainText,key)
	fmt.Println("====", string(cipherByte))
	decText := AesDencryptCTR(cipherByte, key)
	fmt.Println("解密====", string(decText))
}