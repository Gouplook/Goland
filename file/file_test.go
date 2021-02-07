/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/7 13:43
@Description: 关于文件测试用例

*********************************************/
package file

import (
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
}

// 测试io文件拷贝工作
func TestIoFileCopy(t *testing.T) {
	dstName := "fileread.txt"
	srcName := "filewrite.txt"
	IoFileCopy(dstName , srcName)
}

// 测试ioutil写文件
func TestIoutilFileWrite(t *testing.T) {
	IoutilFileWrite("filewrite.txt")
}