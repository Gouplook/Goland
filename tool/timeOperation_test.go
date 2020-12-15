/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/25 上午11:28

*******************************************/
package tool

import (
	"fmt"
	"testing"
	"time"
)

func TestBasicTime(t *testing.T) {
	BasicTime()
}

func TestTimeRange(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	b,e := TimeRange(now)
	fmt.Println("beginTime:= ",b)
	fmt.Println("endTime:= ",e)
}