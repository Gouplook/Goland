/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/4 16:02
@Description:

*********************************************/
package gochannel

import (
	"fmt"
	"time"
)

// 写数据
func WriteData(intChan chan int) {
	i := 0
	//for i := 1; i <= 10; i++ {
	//放入数据
	intChan <- i + 1
	//fmt.Println("writeData ", )
	//}
	close(intChan) //关闭
}

// 读数据
func ReadData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)

	}
	//readData 读取完数据后，通知主线程
	time.Sleep(time.Second)

	exitChan <- true

	close(exitChan)
}
