/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/25 上午11:26

*******************************************/
package tool

import (
	"fmt"
	"time"
)

/*
	函数：
	Now() Time   当前Time
	Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time //返回一个设置的time类型
	Since(t Time) Duration //time.Now().Sub(t)
	Unix(sec int64, nsec int64) Time // 时间戳转时间 1sec = 1nsec * 1e6 , sec 10位时间戳

	方法：
	(t Time) Add(d Duration) Time // returns the time t+d.
	(t Time) AddDate(years int, months int, days int) Time
	(t Time) Sub(u Time) Duration   //计算时间差
	(t Time) Unix() int64  10位时间戳
	(t Time) UnixNano() int64 16位时间戳
	(t Time) Equal(u Time) bool // 比较两个time相等
	(t Time) After(u Time) bool // reports whether the time instant t is after u.
	(t Time) Before(u Time) bool // reports whether the time instant t is before u.
	(t Time) IsZero() bool  // reports whether t represents the zero time instant, January 1, year 1, 00:00:00 UTC.
	(t Time) UTC() Time // returns t with the location set to UTC.
	(t Time) Local() Time // returns t with the location set to local time.
	(t Time) In(loc *Location) Time //设置为指定location
	(t Time) Location() *Location // returns the time zone information associated with t.
	(t Time) Zone() (name string, offset int) // name of the zone (such as "CET") and its offset in seconds east of UTC.

	// 获取当天零点时间戳
    timeStr := time.Now().Format("2006-01-02")


*/

// 时间相关总结
func BasicTime() {
	// Timestamp 时间戳转时间
	now := time.Now()
	local := time.Now().Local()
	timestmap := time.Now().Local().Unix()
	localFroml := time.Now().Local().Format("2006-01-02") // time --> string

	// string -> time
	strToTime, _ := time.Parse("2006-01-02", localFroml)
	//1606875723  将时间字符串转换为时间戳
	stamp, _ := time.ParseInLocation("2006-01-02", "2020-12-02",time.Local)


	fmt.Println("now time: ", now)
	fmt.Println("local time: ", local)
	fmt.Println("timestmap: ", timestmap)
	fmt.Println("localFroml: ", localFroml)
	fmt.Println("strToTime: ", strToTime)
	fmt.Println("stamp",stamp.Unix())
}

// 获取当天时间段 ：2020-12-14 00:00:00~2020-12-14 23:59:59
func TimeRange(now time.Time) (bTime, eTime time.Time) {
	local, _ := time.LoadLocation("Asia/Shanghai")
	bTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, local) // 2020-12-14 00:00:00
	eTime = bTime.AddDate(0, 0, 1).Add(-1 * time.Second) // 2020-12-14 23:59:59
	return
}


//字符串转化为时间戳
//@param  string timeStr 日期字符串
//@return int64
func StrtoTime(timeStr string, timelayouts... string) int64 {
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	if len(timelayouts) > 0 {
		timeLayout = timelayouts[0]
	}
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, timeStr, loc) //使用模板在对应时区转化为time.time类型
	return  theTime.Unix()
}

//时间戳转化为字符串
//@param  int64  timestamp  时间戳
//@return string
func TimeToStr(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006/01/02 15:04:05")
}