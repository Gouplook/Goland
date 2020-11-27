/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/24 下午2:12

*******************************************/
package beegoconf

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func BeeGoConf(){


}
// beego 默认解析ini 配置文件   key通过 section::key 的方式获取
func GetIniFileConf(){
	iniConf, _ := config.NewConfig("ini", "../ini/conf.ini")
	enabled := iniConf.String("Section::enabled")
	fmt.Println(enabled)
}


//
//func GetYamlFileConf(){
//	iniConf, _ := config.NewConfig("yaml", "../yaml/conf.yaml")
//	host := iniConf.String("host")
//
//	fmt.Println(host)
//}


