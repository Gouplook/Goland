/*
 * @Author: your name
 * @Date: 2021-03-08 13:06:11
 * @LastEditTime: 2021-03-08 16:42:25
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
}
