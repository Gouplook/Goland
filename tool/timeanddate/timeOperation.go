/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/25 上午11:26

*******************************************/
package timeanddate

import (
	"fmt"
	"regexp"
	"strconv"
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
	nowForml := now.Format("2006-01")

	// string -> time
	strToTime, _ := time.Parse("2006-01-02", localFroml)
	//1606875723  将时间字符串转换为时间戳
	stamp, _ := time.ParseInLocation("2006-01-02", "2020-12-02",time.Local)


	fmt.Println("now time: ", now)
	fmt.Println("local time: ", local)
	fmt.Println("timestmap: ", timestmap)
	fmt.Println("localFroml: ", localFroml)
	fmt.Println("nowForml: ", nowForml)
	fmt.Println("strToTime: ", strToTime)
	fmt.Println("stamp",stamp.Unix())  // 转化成int
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
	timeLayout := "2006-01-02 15:04:05"                            // 转化所需模板
	if len(timelayouts) > 0 {
		timeLayout = timelayouts[0]
	}
	loc, _ := time.LoadLocation("Local")                    // 重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, timeStr, loc) //  使用模板在对应时区转化为time.time类型
	return  theTime.Unix()
}

//时间戳转化为字符串
//@param  int64  timestamp  时间戳
//@return string
func TimeToStr(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

//时间戳转 数据库的Date格式
func TimeToDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02")
}

// 字符串时间格式转time.Time
func StrToTimeTime(strTime string, format string) (t time.Time, err error){
	local,_ := time.LoadLocation("Asia/Shanghai")
	t,err = time.ParseInLocation(format,strTime,local)
	if err != nil {
		panic(err)
		return
	}
	return t,err
}


// 根据身份证号---获取生日
func GetBirthday(str string) string {
	reg, err := regexp.Compile("^(\\d{6})(\\d{8})(.*)")
	if err != nil {
		return ""
	}

	if reg.MatchString(str) == true {
		submatch := reg.FindStringSubmatch(str)

		test, _ := time.Parse("20060102", submatch[2])
		return test.Format("2006-01-02")
	}
	return ""
}
//根据身份证号---获取年龄
func GetAge(date string) (age int) {
	d, _ := time.Parse("2006-01-02", date)
	year := d.Year()
	if year <= 0 {
		age = -1
	}

	nowyear := time.Now().Year()
	age = nowyear - year
	return
}


// 获取年份中的月份。
func YearMonth(year string)(start,end string) {
	lastDate := time.Now()
	// 判断选择的年份是否小于当前年份
	if year < strconv.Itoa(lastDate.Year()) && year != "" {
		start = year + "-" + "01"
		end = year + "-" + "12"
		return
	}
	// 默认参数
	lastDate.AddDate(0,-1,0)
	start = strconv.Itoa(lastDate.Year()) + "-" + "01"
	end = lastDate.Format("2006-01")

	return
}

//  时间转化  当日有效 , 返回起终时间戳
func DataEffects (date int ) (startTime, endTime int ){
	loc,_  := time.LoadLocation("Local")
	now := time.Now().Format("2006-01-02")
	firstTime ,_ := time.ParseInLocation("2006-01-02",now,loc)
	startTime = int(firstTime.Unix())
	endTime = startTime+ (86400 *date)
	return
}

// 次日生效
func DataMorrowEffects (date int ) (startTime, endTime int ){

	loc,_  := time.LoadLocation("Local")
	tomorrow := time.Now().Unix()
	tomorrow += 86400
	now := time.Unix(tomorrow,0).Format("2006-01-02")
	firstTime ,_ := time.ParseInLocation("2006-01-02",now,loc)
	startTime = int(firstTime.Unix())
	endTime = startTime+ (86400 *date)
	return
}


// 判读时间是否在给定时间内 t为时间戳
func DataAfter(t int) {
	times := time.Now()
	timeTime := time.Unix(int64(t),0)
	// 判读t 传过来的时间是否在当前时间之后。
	if times.After(timeTime) {
		fmt.Println("在当前时间之后.....")
	}
}

