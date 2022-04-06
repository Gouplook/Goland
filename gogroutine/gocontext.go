package gogroutine

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// 控制并发goroutine
func Rpc(ctx context.Context, url string) error {
	result := make(chan int)
	err := make(chan error)

	go func() {
		// 进行RPC调用，并且返回是否成功，成功通过result传递成功信息，错误通过error传递错误信息
		isSuccess := true
		if isSuccess {
			result <- 1
		} else {
			err <- errors.New("some error happen")
		}
	}()

	select {
	case <-ctx.Done():
		// 其他RPC调用调用失败
		fmt.Println(url)
		return ctx.Err()
	case e := <-err:
		// 本RPC调用失败，返回错误信息
		fmt.Println(url)
		return e
	case <-result:
		// 本RPC调用成功，不返回错误信息
		return nil
	}
}

// 设置时间停止goroutine
func DealContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("子携程设置5s结束....")
			return
		default:
			fmt.Println("没有启动子携程....")
			time.Sleep(time.Second * 1)
		}
	}
}

// context 传替value
func ContextValue(ctx context.Context) {
	for {
		select {
		default:
			// 执行任务，并获取主携程的值
			fmt.Println("STOP goroutine and get value", ctx.Value("key"))
			return
		}
	}
}
