/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/7 17:25
@Description:

*********************************************/
package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// bufio 对io方法进行封装

// io ReadString : 是将外部文件读到（应用程序中）输出流中
func IoFileRead(name string) {
	file, _ := os.Open(name)
	reader := bufio.NewReader(file)
	defer file.Close()
	// 每次读以delim 作为分隔符字符串
	for {
		str, err := reader.ReadString('\n')
		// var EOF = errors.New("EOF")
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
}

func IoFileWrite(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("文件创建失败")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	n, err := writer.WriteString("func LimitReader(r Reader, n int64) Reader11209871")
	// 注意：
	// 因为writer是带缓存，因此在调用WriterString方法时，其实
	// 内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	// 真正写入到文件中， 否则文件中会没有数据!!!
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

// 利用io.Copy文件拷贝
func IoFileCopy(dstName, srcName string) {
	// open  openFile 区别
	//  os.Open 打开的文件只有读权限，
	//  os.OpenFile 可以添加写权限
	srcFile, _ := os.Open(srcName)
	defer srcFile.Close()
	srcRead := bufio.NewReader(srcFile)

	dstFile, _ := os.OpenFile(dstName, os.O_RDWR|os.O_APPEND, 0777)
	defer dstFile.Close()
	dstWrite := bufio.NewWriter(dstFile)

	io.Copy(dstWrite, srcRead)

}

func IoutilFileRead(name string) {
	file, err := os.OpenFile(name, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 一次性读取所有文件
	by, _ := ioutil.ReadAll(file)
	fmt.Println(string(by))

}
func IoutilFileWrite(name string) {
	data := []byte("大学之道，在明明德，在亲民，止于至善")
	err := ioutil.WriteFile(name, data, 0777)
	if err != nil {
		panic(err)
	}
}
