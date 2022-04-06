package gogroutine

import (
	"fmt"
	"sync"
	"time"
)

// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
// 最后显示出来。要求使用goroutine完成
// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
// 2. 我们启动的协程多个，统计的将结果放入到 map中
// 3. map 应该做出一个全局的.

var FactorialMap = make(map[int]int, 100)
var Lock sync.Mutex

func Factorial(n int, wg *sync.WaitGroup) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	Lock.Lock()
	FactorialMap[n] = res
	Lock.Unlock()
	wg.Done()
}

// 读写锁
// 读锁会阻塞写入数据，不影响其他goroutine读操作。（写同理）
func ReadLock(lock *sync.RWMutex) {
	lock.RLock()
	fmt.Println("read data....")
	time.Sleep(time.Second * 1)
	defer lock.RUnlock()

}

func WriteLock(lock *sync.RWMutex) {
	lock.Lock()
	fmt.Println("write data....")
	time.Sleep(time.Second * 1)
	defer lock.Unlock()
}

// sync.Map
func SyncMapWrite(mp *sync.Map, wg *sync.WaitGroup) {
	fmt.Println("syncMap Write .....")
	mp.Store("name", "DZ001")
	mp.Store("name2", "DZ002")
	mp.Store("nam3", "DZ003")
	//for i := 1; i < 5; i++ {
	//	mp.Store(i, i*i)
	//}
	defer wg.Done()
}
func SyncMapRead(mp *sync.Map, wg *sync.WaitGroup) {

	fmt.Printf("syncMap Read.....")
	//for i := 1; i < 5; i++ {
	//	fmt.Println(mp.Load(i))
	//}
	mp.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true

	})
	defer wg.Done()
}
