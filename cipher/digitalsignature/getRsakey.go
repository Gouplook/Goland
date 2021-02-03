/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/3 16:12
@Description:

*********************************************/
package digitalsignature

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// 生成私钥流程
// 1：使用rsa中的GenerateKey方法生成私钥
// 2：通过X509标准得到的RAS私钥序列化特定的字符串ASN.1 PKCS#1 DER编码字符串
// 3：将私钥串设置到pem格式块中
// 4：通过pem将设置好的数据进行编码，并写入磁盘文件中

// 生成公钥流程
// 1：从得到私钥的对象中将公钥信息取出
// 2：通过X509标准将得到的RSA公钥序列化为字符串
// 3：将公钥字符串设置到pem格式块中
// 4：通过pem将设置好的数据进行编码, 并写入磁盘文件

// 生成私钥
// 注意事项：
func GetRsakey(bits int) {
	// 1: 使用rsa中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
		return
	}
	// 2：通过X509标准得到的RAS私钥序列化特定的字符串ASN.1 PKCS#1 DER编码字符串
	privateByte := x509.MarshalPKCS1PrivateKey(privateKey)

	// 3：将私钥串设置到pem格式块中
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateByte,
	}

	// 4: 组装好的pem块进行编码，写如磁盘中
	file, err := os.Create("private.pem")
	if err != nil {
		panic(err)
		return
	}
	err = pem.Encode(file, block)
	if err != nil {
		panic(err)
		return
	}
	file.Close()

	// ========================= 公钥 ================================
	// 从私钥中取出公钥（由于私钥结构中包含公钥)
	// 1：从得到私钥的对象中将公钥信息取出
	publicKey := privateKey.PublicKey

	// 2：通过X509标准将得到的RSA公钥序列化为字符串
	// ******* 注意问题，公钥加密方法，和解密使用的方法要一致。********
	//publicByte := x509.MarshalPKCS1PublicKey(&publicKey)
	publicByte,err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
		return
	}

	// 3：将公钥字符串设置到pem格式块中
	publicBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicByte,
	}

	// 4：通过pem将设置好的数据进行编码, 并写入磁盘文件
	publicFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
		return
	}

	err = pem.Encode(publicFile, publicBlock)
	if err != nil {
		panic(err)
		return
	}
	publicFile.Close()

}