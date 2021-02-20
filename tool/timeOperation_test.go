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

func TestStrtoTimeTime(t *testing.T) {
	//字符串转化为时间戳

	timesp := StrtoTime("2021-02-19","2006-01-02")
	fmt.Println(timesp)
	// 字符串时间格式转time.Time
	tm, _ := StrToTimeTime("2001-02-03","2006-01-02")
	fmt.Println(tm)
}