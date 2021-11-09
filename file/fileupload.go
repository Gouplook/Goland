/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/20 13:47
@Description:  文件上传与下载

*********************************************/
package file

import (
	"GoInduction/functions"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type ReplyFileInfo struct {
	Id   int
	Hash string
	Path string
}
// 根据url地址获取远程图片
// 1：根据url下载到制定位置
// 2：获取图片信息保存到数据库中
func SaveImgFromUrl (imgUrl string)(reply ReplyFileInfo,err error) {
	// 1: 检查是否是url地址  "http://"  "https://"

	// 2: 下载图片  安全传输层协议（TLS）用于在两个通信应用程序之间提供保密性和数据完整性
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{
		Transport: tr,
		Timeout: time.Second *30,  // 30s
	}
	request := http.Request{}
	request.Method = http.MethodGet
	// 将srtURl解析成URL结构
	parse,err := url.Parse(imgUrl)
	if err != nil {
		return
	}
	request.URL = parse
	response, err := client.Do(&request)
	defer response.Body.Close()
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		return
	}
	body := response.Body

	// 3: 对请求数据进行加密，采用md5
	imgByte,_  := ioutil.ReadAll(body)
	objMd5 := md5.New()
	objMd5.Write(imgByte)
	// 利用Hex将数据src编码为字符串
	strMd5 := hex.EncodeToString(objMd5.Sum(nil))

	// 4: 对请求数据进行sha256加密得到strSha256字符串
	objSha256 := sha256.New()
	objSha256.Write(imgByte)
	strSha256 := hex.EncodeToString(objSha256.Sum(nil))

	// 获取图片大小 imgSize
	imgSize,_ := strconv.ParseInt(response.Header.Get("Content-Length"),10,64)

	// 5：利用获取的strMd5和strSha256字符串，数据库业务操作（查重）
	fmt.Println(strMd5)
	fmt.Println(strSha256)

	// 如何获取hash字符串值,通过go.uid框架
	u := uuid.NewV4()
	// 9d2b88b9-3648-4329-b7ab-f838be944b74
	// 9d2/b88/b93/648/432/9b7/abf838be944b74
	hash := strings.Replace(u.String(),"-","", -1)
	filename := functions.GetFileNameByHash(hash)
	path := filepath.Join("upload/image",filename)
	// 创建一个目录 给权限0777
	err = os.MkdirAll(filepath.Dir(path),os.ModePerm)
	if err != nil {
		return
	}
	f,_ := os.OpenFile(path,os.O_WRONLY | os.O_CREATE,0644)
	defer  f.Close()
	// 保存文件
	_, err = f.Write(imgByte)

	// 获取图片的名称和后缀
	fileNames := strings.Split(imgUrl,"/")
	fileExt := response.Header.Get("Content-Type")
	if len(fileExt)==0 {
		fileExt="png"
	}
	fmt.Println(fileNames)
	fmt.Println(imgSize)
	// 将文件信息保存到数据库中
	// 需保存字段 name,ext,size, type, strMd5,strSha256
	return
}

// 根据hashs查图片