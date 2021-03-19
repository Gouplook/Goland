/*
 * @Author: your name
 * @Date: 2021-03-08 13:06:11
 * @LastEditTime: 2021-03-09 13:39:23
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Goland/conf/ini/ini.go
 */
/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/24 下午1:36

*******************************************/
package ini

import (
	"fmt"
	"log"

	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego/config"
	"gopkg.in/gcfg.v1"
)

type Section struct {
	Enabled bool
	Path    string
}

func IniReadConf() {
	// 反序列化
	conf := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&conf, "conf.ini")
	if err != nil {
		fmt.Println("Failed to parse config file")
	}
	fmt.Println(conf.Section.Enabled)
	fmt.Println(conf.Section.Path)
}

// go get github.com/Unknwon/goconfig  直接读取
func GoConfigReadConf() {
	// 直接获取
	cfg, _ := goconfig.LoadConfigFile("conf.ini")
	enabled, _ := cfg.GetValue("Section", "enabled")

	path, _ := cfg.GetValue("Section", "path") // 结构体区分大小写，Section，

	fmt.Println(enabled)
	fmt.Println(path)
	// 同时也获取注释的内容
	path = cfg.GetKeyComments("Section", "path")
	path1, _ := cfg.GetValue("Section", "path")
	dbType, _ := cfg.GetValue("db", "type")

	fmt.Println(enabled)
	fmt.Println(path)
	fmt.Println("path1 = ", path1)
	fmt.Println("dbType= ", dbType)
}

// 利用beego config 读取配置文件,beego 默认解析ini 配置文件
// key通过 section::key 的方式获取
func Beegoconfig() {
	// 初始化解析对象
	inc, err := config.NewConfig("ini", "conf.ini")
	if err != nil {
		log.Fatal("config NewConfig error")
		return
	}
	// 调用配置方法
	dbtype := inc.String("db::type")

	fmt.Println(dbtype)

}
