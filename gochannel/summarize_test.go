/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/4 16:43
@Description:

*********************************************/
package gochannel

import (
	"fmt"
	"testing"
	"time"
)

func TestCpuNum(t *testing.T) {
	CpuNum()
}

func TestSelectChannel(t *testing.T) {
	SelectChannel()
}

// 测试阶乘问题
func TestFactorial(t *testing.T) {
	// 开启携程gogroutine
	//for i := 1; i < 10; i++ {
	//	go Factorial(i)
	//}
	//lock.Lock()
	//for i, v := range FactorialMap {
	//	fmt.Printf("map[%d]=%d\n", i, v)
	//}
	//lock.Unlock()

}

// 测试素数
func TestPrimeNum(t *testing.T) {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000) //放入结果
	// 标识退出的管道
	exitChan := make(chan bool, 8) // 4个

	start := time.Now().Unix()
	// 开启一个协程，向 intChan放入 1-8000个数'
	go PutNum(intChan, 100)
	// 开启8个协程，从 intChan取出数据，并判断是否为素数.
	for i := 0; i < 8; i++ {
		go PrimeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)
		close(primeChan)
	}()

	//遍历我们的 primeChan ,把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}

}
