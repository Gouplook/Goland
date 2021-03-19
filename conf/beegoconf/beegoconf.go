/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/24 下午2:12

*******************************************/
package beegoconf

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/config"
)

// 待完善
func BeeGoConf() {
	cg, err := config.NewConfig("ini", "app.conf")

	if err != nil {
		log.Fatal("config.NewConfig error")
		return
	}
	redisHost := cg.String("db.maxconn")
	fmt.Println(redisHost)

}
