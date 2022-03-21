/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/4 16:03
@Description:

*********************************************/
package gochannel

import (
	"fmt"
	"testing"
)

func TestReadData(t *testing.T) {
	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go WriteData(intChan)
	go ReadData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

	fmt.Println()

}
