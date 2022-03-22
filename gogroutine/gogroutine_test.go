package gogroutine

import (
	"sync"
	"testing"
)

func TestNumcpu(t *testing.T) {
	Numcpu()
}

var (
	mapDate = make(map[int]int, 10)
	lock    sync.Mutex
)

func Ggogroutine(n int) {

}
