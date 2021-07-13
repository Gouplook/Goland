/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/3 15:03
@Description: RSA数字签名

*********************************************/
package digitalsignature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// RSA签名 = plainText进行哈希运算值 + 私钥
// 1： 将原始数据对其进行哈希运算  ----->   散列值
// 2： 使用非对称加密的 私钥 对散列值加密 -> 签名
// 3： 将原始数据和签名一并发送给对方

// 验签
// Data = 原始数据 + 数字签名
// 数字签名, 需要使用 公钥 进行解密-----> 得到散列值

// RSA签名 - (利用私钥进行签名）
func SignatureRSA(plainText []byte, fileName string) []byte {
	//1. 打开磁盘的私钥文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}


	//2. 将私钥文件中的内容读出
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	defer file.Close()

	//3. 使用pem对私钥数据解码, 得到了pem.Block结构体变量
	block, _ := pem.Decode(buf)

	//4. x509将数据解析成私钥结构体 -> 得到了私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	//5. 创建一个哈希对象 -> md5/sha1 -> sha512
	rsaHash := sha512.New()

	//6. 给哈希对象添加数据
	rsaHash.Write(plainText)

	//7. 计算哈希值(得到是一个固定长度的哈希值）
	hashText := rsaHash.Sum(nil)

	//8. 使用rsa中的函数对散列值签名
	signText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashText)
	if err != nil {
		panic(err)
	}

	return signText
}

// RSA签名验证
// @sigText 签名字符串
// @publicKey 公钥
func VerifyRSA(plainText, sigText []byte, publicKeyFile string) bool {
	//1. 打开公钥文件, 将文件内容读出 - []byte
	file, err := os.Open(publicKeyFile)
	if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat()
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	defer file.Close()

	//2. 使用pem解码 -> 得到pem.Block结构体变量
	block, _ := pem.Decode(buf)

	//3. 使用x509对pem.Block中的Bytes变量中的数据进行解析 ->  得到一接口
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	//4. 进行类型断言 -> 得到了公钥结构体
	publicKey := pubInterface.(*rsa.PublicKey)

	//5. 对原始消息进行哈希运算(和签名使用的哈希算法一致) -> 散列值
	hashText := sha512.Sum512(plainText) // 返回[64]byte

	//6. 签名认证 - rsa中的函数
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashText[:], sigText)
	if err == nil {
		return true
	}

	return false
}
