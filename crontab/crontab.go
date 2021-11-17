/**
 * @Author: yinjinlin
 * @File:  crontab
 * @Description:
 * @Date: 2021/10/21 下午2:19
 */

package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

// 每5s启动一次
func main(){
	c := cron.New()
	// 每5s触发一次
	spec := " */5 * * * * ?"  // 秒-分-时-日-月-周【忽略】
	_ = c.AddFunc(spec,printf)
	c.Start()
	select {}

}

func printf(){
	local, _ := time.LoadLocation("Local")
	localTime := time.Now().In(local).Unix()
	localTimeStr := time.Unix(localTime,0).Format("2006-01-02-15-04-05")
	fmt.Println(localTimeStr)
}



