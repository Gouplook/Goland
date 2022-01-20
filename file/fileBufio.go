/**
 * @Author: yinjinlin
 * @File:  fileBufio.go
 * @Description:
 * @Date: 2022/1/20 下午4:47
 */

package file

import (
	"bufio"
	"fmt"
	"os"
)

func FileBufio(){
	file,_  := os.Open("./fileread.txt")
	inputReader := bufio.NewReader(file)

	defer file.Close()

	fmt.Println("Please enter some input:....")
	// 按下回车键后停止
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
