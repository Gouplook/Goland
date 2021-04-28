/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/22 17:18
@Description:

*********************************************/
package goroutine

import (
	"fmt"
	"sync"
)

func GoWait() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		//go func(i int ) {
		//	defer wg.Done()
		//	fmt.Println("i= ",i)
		//}(i)
		go FunTest(i, &wg)
	}
	wg.Wait()
	fmt.Println("最后输出......")

}

func FunTest(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("i= ", i)
}
