/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/1 14:05
@Description:

*********************************************/
package sysmmertic

import (
	"fmt"
	"testing"
)



func TestDesDecrpt(t *testing.T) {

	plainText := "创建一个使用cbc分组接口"

	key := "12345678"
	// 加密
	fmt.Printf("加密前：%v", plainText )
	fmt.Println()
	cipherByte := DesEncrypt([]byte(plainText),[]byte(key))
	fmt.Println("====", string(cipherByte))


	// 解密
	decText := DesDecrpt(cipherByte,[]byte(key))
	fmt.Println("解密====", string(decText))


}