/*
 * @Author: your name
 * @Date: 2021-03-08 13:06:11
 * @LastEditTime: 2021-03-23 13:53:53
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Goland/file/file_test.go
 */

package file

import (
	"fmt"
	"testing"
)

// 测试文件输出流
func TestFileStream(t *testing.T) {
	FileStream("fileread.txt")
	fmt.Println(5 + 13)
	fmt.Println("")
	fmt.Println(300000*0.3*12 *0.2)
}

// 测试 io 文件读写
func TestIoFileReadWrite(t *testing.T) {
	IoFileRead("fileread.txt")
	name := "filewrite.txt"
	IoFileWrite(name)
}

// 测试ioutil 读文件
func TestIoutilFile(t *testing.T) {
	IoutilFileRead("fileread.txt")
	fmt.Println("TestIoutilFile22233")
	// t.Log("xxxue443")
}

// 测试io文件拷贝工作
func TestIoFileCopy(t *testing.T) {
	dstName := "fileread.txt"
	srcName := "filewrite.txt"
	IoFileCopy(dstName, srcName)
}

// 测试ioutil写文件
func TestIoutilFileWrite(t *testing.T) {
	IoutilFileWrite("filewrite.txt")
	fmt.Println("hello world.......")
	t.Log("122222....")
}
