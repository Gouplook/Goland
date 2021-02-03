/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/1 15:07
@Description: 对称加密/解密 密钥 是相同的

*********************************************/
package sysmmertic

import (
	"crypto/aes"
	"crypto/cipher"
)

var iv = []byte("1234567890123456")
// 分组模式
// OFB - 不推荐使用
// CFB - 不推荐使用
// CTR - 推荐使用

// ------------------------cbc-----------------------------------
// 加密 模式 cbc
func AesEncrypt(plainText, key []byte) []byte {
	// 1. 建一个底层使用aes的密码接口,参数key为密钥，长度只能是16、24、32字节
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
		return nil
	}

	// 2. 如果使用的是cbc/ecb分组模式需要对明文分组进行填充
	newPlainText := paddingLastGroup(plainText, block.BlockSize())

	// 3. 创建一个密码分组模式的接口对象
	//iv := "12345678abcdefgh"
	//iv := plainText[:block.BlockSize()]
	blockMode := cipher.NewCBCEncrypter(block,iv)

	// 4.  加密
	dst := make([]byte, len(newPlainText))
	src := newPlainText
	blockMode.CryptBlocks(dst, src)

	return dst
}

// 解密
func AesDencrypt(cipherText, key []byte) []byte {
	// 1. 建立一个底层使用AES的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
		return nil
	}
	//
	//iv := []byte("12345678abcdefgh")
	blockMode := cipher.NewCBCDecrypter(block,iv)

	dst := make([]byte,len(cipherText))
	src := cipherText
	blockMode.CryptBlocks(dst, src)

	// 去掉填充的尾部
	plainText := unPaddingLastGroup(dst)
	return plainText
}


//--------------------OFB 不推荐使用-----------------------
// 该模式下加密，不需要填充
// 加密 OFB分组模式 加密和解密相同，相当于两次异或（XORKeyStream）
func AesEncryptOfb(plainText, key []byte) []byte{
	// 1. 建一个底层使用aes的密码接口
	block,err := aes.NewCipher(key)
	if err != nil {
		panic(err)
		return nil
	}

	// 2. 创建一个密码分组模式的接口对象
	stream := cipher.NewOFB(block, iv)
	dst := make([]byte,len(plainText))
	src := plainText

	// 对接口对象进行加解密操作
	stream.XORKeyStream(dst, src )
	return dst
}


// 解密
func AesDencryptOfb(ciphertext, key []byte) []byte {
	// 1: 创建一个底层的aes接口
	block,err := aes.NewCipher(key)
	if err != nil {
		panic(err)
		return nil
	}

	// 2: 创建和一个密码分组模式的接口对象
	stream := cipher.NewOFB(block, iv)
	dst := make([]byte,len(ciphertext))
	src := ciphertext

	// 3: 对接口对象进行加解密操作
	stream.XORKeyStream(dst,src)
	return dst

}

// -------------------------CTR----------------
func AesEncryptCTR(plainText,key []byte) []byte{
	block,err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	stream := cipher.NewCTR(block,iv)
	src := plainText
	dst := make([]byte,len(plainText))

	stream.XORKeyStream(dst, src)

	return dst
}

func AesDencryptCTR(cipherText,key []byte)[]byte {
	block,err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	stream := cipher.NewCTR(block,iv)
	src := cipherText
	dst := make([]byte,len(cipherText))
	stream.XORKeyStream(dst,src)
	return dst
}


