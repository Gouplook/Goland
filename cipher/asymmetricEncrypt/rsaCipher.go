/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/2 14:03
@Description: RSA非对称加密

*********************************************/
package asymmetricEncrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
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
	// 通过pem.Encode  写入到磁盘中
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
	publicByte := x509.MarshalPKCS1PublicKey(&publicKey)
	//publicByte,err := x509.MarshalPKIXPublicKey(&publicKey)
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

// RSA加密：
// 1：将公钥文件中的公钥读出，得到使用pem编码的字符串
// 2：将得到的字符串进行解码
// 3：使用x509将编码之后的公钥解析出来
// 4：使用得到的公钥通过rsa进行数据加密

// RSA 解密
// 1：将私钥文件中的私钥读出来，得到使用pem编码的字符串
// 2：将得到的字符串解码
// 3：使用x509将编码之后的私钥解析出来
// 4：使用得到的私钥通过RSA进行数据解密

// RAS加密  采用公钥加密
//@fileName  公钥或私钥 存放的文件路径
func RSAEncrypt(plainText []byte, fileName string) []byte {
	// 1：将公钥文件中的公钥读出，得到使用pem编码的字符串
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	buf := make([]byte, fileInfo.Size())

	file.Read(buf)
	defer file.Close()
	// 2：将得到的字符串进行解码
	block, _ := pem.Decode(buf)

	// 3：通过X509标准将得到的RSA公钥序列化为字符串
	pubInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)  // 密钥解析
	//pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	// 断言类型转换
	//pubKey := pubInterface.(*rsa.PublicKey)
	// 4：使用得到的公钥通过rsa进行数据加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface, plainText)

	if err != nil {
		panic(err)
	}

	return cipherText
}

// RSA 解密 私钥解密
//@fileName	：私钥/公钥 文件
func RSADencrpt(cipherText []byte, fileName string) []byte {
	// 1: 加载
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()

	block, _ := pem.Decode(buf)
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, cipherText)
	if err != nil {
		panic(err)
	}

	return plainText
}
 // Hash 加密
func AsyHash() {
	// 1: 创建哈希接口对象
	hash := sha256.New()

	// 2. 添加数据
	src :=[]byte("哈希接口对象加密")
	hash.Write(src)
	hash.Write(src)

	// 3. 计算结果
	res := hash.Sum(nil)

	// 4. 格式化为16进制形式
	str := hex.EncodeToString(res)
	src2 := hex.EncodeToString(res)
	fmt.Println(src2)
	fmt.Printf("%s\n", str)
}
