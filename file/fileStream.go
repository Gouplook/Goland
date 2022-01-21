/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/7 13:42
@Description: 文件输出流和输入流

*********************************************/
package file

import (
	"fmt"
	"os"
)


// 输出流： 将应用程序中数据写如到外部资源中
func FileStream(name string){
	// 1：读取文件
	// 构建字节切片的时候如果需要读取文件,不要使用os.Open,该方法获取的文件只能读取，无法写入
	// name:文件名，flag:操作方， perm:文件权限
	// fileName 文件
	file, err := os.OpenFile(name,os.O_RDWR , 0777)

	if err != nil {
		fmt.Println("文件不存在，正在创建....")
		file,_ = os.Create("fileread.txt")
	}
	defer file.Close()

	// 向输出流中写入数据
	_ ,_ = file.Write([]byte("要写入的内容1\\r\\n要写入的内容2\\r\\n\\t要写入的内容3\\r\\n"))
	_ ,_= file.Write([]byte("\n"))
	_ ,_ = file.WriteString("==写入这一行......")
	_ ,_ = file.WriteAt([]byte("HHHHHHH"),2)   // 此方法与追加不用混用。

}




