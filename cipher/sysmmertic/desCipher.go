/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/1 09:53
@Description: 对称加密 Des 加密方式

*********************************************/
package sysmmertic

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

// 增加填充数据
// 编写填充函数, 如果最后一个分组字节数不够, 填充
// 若字节数刚好合适，添加一个新的分组；
// 填充个的字节的值 == 缺少的字节的数
func paddingLastGroup(plainText []byte, bloclSize int) []byte{
	padNum := bloclSize - len(plainText) % bloclSize
	ch := []byte{
		byte(padNum),
	}

	newText := bytes.Repeat(ch, padNum)
	newText = append(plainText, newText...)
	return  newText
}

// 去掉填充的数据
func unPaddingLastGroup(plainText []byte) []byte{
	leng := len(plainText)
	lastChar := plainText[leng -1]
	number := int(lastChar)
	return plainText[:leng - number]
}

// des加密
func DesEncrypt(plainText, key []byte) []byte {
	// 1. 建一个底层使用des的密码接口
	cipherBlock, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	// 2. 明文填充
	newPlainText := paddingLastGroup(plainText, cipherBlock.BlockSize())

	iv := []byte("12345678")
	// 3: 创建一个使用cbc分组接口
	blockMode := cipher.NewCBCEncrypter(cipherBlock,iv)

	src := newPlainText
	dst := make([]byte, len(newPlainText))
	// 4. 加密
	blockMode.CryptBlocks(dst, src)
	return dst
}
// @cipherText	:加密文本
// @key			：密钥（由于是对称加密，加密和解密密钥 是相同的）
func DesDecrpt(cipherText, key[]byte) []byte{

	// 1.创建一个底层使用des的密码接口
	cipherBlock, err  := des.NewCipher(key)
	if err != nil {
		return nil
	}
	iv := []byte("12345678")  //
	// 2. 创建一个使用的cbc模式解密的接口
	stream := cipher.NewCBCDecrypter(cipherBlock, iv)

	src := cipherText
	dst := make([]byte,len(cipherText))

	// 3. 解密
	stream.CryptBlocks(dst,src)

	// 4. cipherText 现存储的是明文，需要删除填充的数据
	plainText := unPaddingLastGroup(dst)
	return plainText
}