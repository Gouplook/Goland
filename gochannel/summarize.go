/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/4 16:14
@Description: 管道总结

*********************************************/
package gochannel

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 管道可以声明为只读或者只写
//1. 在默认情况下下，管道是双向
//2  声明为只写
//  var chanWrite chan<- int

//3. 声明为只读
//var chanReade <- chan int


func SelectChannel() {
	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	// 问题，在实际开发中，可能我们不好确定什么关闭该管道.
	// 可以使用select 方式可以解决
	// label:
	for {
		select {
		//注意: 这里，如果intChan一直没有关闭，不会一直阻塞而deadlock
		//，会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("都取不到了，不玩了, 程序员可以加入逻辑\n")
			time.Sleep(time.Second)
			return
		}
	}

}

func CpuNum() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)

	fmt.Println("ok")
}

// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
// 最后显示出来。要求使用goroutine完成
// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
// 2. 我们启动的协程多个，统计的将结果放入到 map中
// 3. map 应该做出一个全局的.

var FactorialMap = make(map[int]int, 10)
var lock sync.Mutex

func Factorial(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	FactorialMap[n] = res
	lock.Unlock()
}

// ------计算素数--------
// 向intChan 放入num个数
func PutNum(intChan chan int, num int) {
	for i := 1; i <= num; i++ {
		intChan <- i
	}
	//关闭intChan
	close(intChan)
}

// 从intChan管道中取出数据，进行判断是否是素数。并放入到primeChan中
func PrimeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true // 假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			//将这个数就放入到primeChan
			primeChan <- num
		}
	}
	//这里不能关闭 primeChan
	//向 exitChan 写入true
	exitChan <- true
}

