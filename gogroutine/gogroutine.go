package gogroutine

import (
	"fmt"
	"runtime"
)

//gogroutine 特点
//1：具有独立栈空间
//2：共享程序堆内存空间
//3：轻量级线程
//4：调度由用户控制

//MPG

//
func Numcpu() {
	numCpu := runtime.NumCPU
	fmt.Println("CPU num := ", numCpu())
}
