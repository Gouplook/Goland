

package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

var (
	HttpClient = &http.Client{
		Timeout: 3*time.Second,
	}
)

// 上传文件
// url          请求地址
// nameField    请求地址上传文件对应field
// fileName     文件名
// file         文件
func uploadFile(url string, params map[string]string, nameField,fileName string, file io.Reader)([]byte, error){

	//
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile(nameField,fileName)
	if err != nil {
		return nil, err
	}
	// 文件内容拷贝到
	_, err = io.Copy(formFile,file)
	if err != nil {
		return nil, err
	}

	// 添加请求参数
	for key,val := range params {
		_ = writer.WriteField(key,val)
	}
	err = writer.Close()
	if err != nil {
		return nil,err
	}

	// 请求
	req, err := http.NewRequest("POST",url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-type", writer.FormDataContentType())
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content,nil
	//


}

