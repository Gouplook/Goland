/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/24 下午2:12

*******************************************/
package beegoconf

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

// 待完善
func BeeGoConf() {

}

// beego 默认解析ini 配置文件   key通过 section::key 的方式获取
func GetIniFileConf() {
	// 初始化一个解析器对象
	iniConf, _ := config.NewConfig("ini", "../ini/conf.ini")
	// 通过对象获取数据
	enabled := iniConf.String("Section::enabled")
	path := iniConf.String("Section::path")
	dbType := iniConf.String("db::type")

	fmt.Println(enabled)
	fmt.Println("Path := ", path)
	fmt.Println("dbType = ",dbType)
	fmt.Println("==========")
	fmt.Println("Section::path")

}
