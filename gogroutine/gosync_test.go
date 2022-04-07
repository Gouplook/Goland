package gogroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestFactorial(t *testing.T) {

	// 等待goroutine结束
	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go Factorial(i, &wg)
	}

	Lock.Lock()
	for k, v := range FactorialMap {
		fmt.Printf("map[%d]= %d", k, v)
		fmt.Println(",")
	}
	Lock.Unlock()

	wg.Wait()

}

func TestRWLock(t *testing.T) {
	l := sync.RWMutex{}

	go WriteLock(&l)
	go WriteLock(&l)
	go ReadLock(&l)
	go ReadLock(&l)
	go ReadLock(&l)
	go ReadLock(&l)

	for {
	}
}

func TestSyncMapRW(t *testing.T) {
	mp := &sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go SyncMapWrite(mp, &wg)
	go SyncMapRead(mp, &wg)

	wg.Wait()

}

// 先查询出来，[]int
// 在根据两个切片进行对比

