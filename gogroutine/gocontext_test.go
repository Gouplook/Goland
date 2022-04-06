package gogroutine

import (
	context2 "context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// context 控制并发
// rpc1 与{rpc2 /3/4 } 并行
func TestRpc(t *testing.T) {
	ctx, cancel := context2.WithCancel(context2.Background())
	// rpc1
	err := Rpc(ctx, "rpc1..")
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	// rpc2
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = Rpc(ctx, "rpc2")
		if err != nil {
			cancel()
			return
		}
	}()
	// RPC3
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = Rpc(ctx, "rpc3")
		if err != nil {
			cancel()
			return
		}
	}()

	// RPC4
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = Rpc(ctx, "rpc4")
		if err != nil {
			cancel()
			return
		}
	}()

	wg.Wait()
}

func TestDealContext(t *testing.T) {
	ctx, _ := context2.WithDeadline(context2.Background(), time.Now().Add(time.Second*5))

	go DealContext(ctx)

	fmt.Println("主携程.....")
	<-time.After(time.Second * 6)

	return

}

func TestContextValue(t *testing.T) {
	ctx := context2.WithValue(context2.Background(), "key", "99802")
	go ContextValue(ctx)

	<-time.After(time.Second * 3)
	return
}
