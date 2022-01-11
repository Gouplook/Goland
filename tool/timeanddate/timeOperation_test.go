/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/25 上午11:28

*******************************************/
package timeanddate

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestBasicTime(t *testing.T) {
	BasicTime()
}

func TestTimeRange(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	b, e := TimeRange(now)
	fmt.Println("beginTime:= ", b)
	fmt.Println("endTime:= ", e)
}

func TestStrtoTimeTime(t *testing.T) {
	// 字符串转化为时间戳

	timesp := StrtoTime("2021-02-19", "2006-01-02")
	fmt.Println(timesp)
	// 字符串时间格式转time.Time
	tm, _ := StrToTimeTime("2001-02-03", "2006-01-02")
	fmt.Println(tm)
}

func TestTimeSoStr(t *testing.T) {
	timesp := "1590561184"
	timesp2, _ := strconv.ParseInt(timesp, 10, 64)
	uTime := TimeToStr(timesp2)

	uTime2 := time.Now().Local().Format("2006-01-02")
	tm := time.Unix(timesp2, 0)

	fmt.Println(tm.Format("2006"))
	fmt.Println(tm.Year())

	fmt.Println(tm.Format("01"))
	fmt.Println(tm.Format("01"))

	fmt.Println(tm.Format("02"))
	fmt.Println(tm.Format("15"))
	fmt.Println(tm.Format("04"))
	fmt.Println(tm.Format("05"))

	fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))
	fmt.Println(uTime2)

	fmt.Println(uTime)
}

func TestGetBirthday(t *testing.T) {
	str := GetBirthday("34082319981024043x")
	fmt.Println(str)
}
func TestGetAge(t *testing.T) {
	str := GetBirthday("34082319981024043x")
	age := GetAge(str)
	fmt.Println(age)
}

//
func TestYearMonth(t *testing.T) {
	strart, end := YearMonth("2021")
	fmt.Println(strart)
	fmt.Println(end)
}
