/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/24 下午1:36

*******************************************/
package ini

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"gopkg.in/gcfg.v1"
)

type Section struct {
	Enabled bool
	Path    string
}

func IniReadConf(){
	// 反序列化
	conf := struct {
		Section struct{
			Enabled bool
			Path    string
		}
	}{}
	err := gcfg.ReadFileInto(&conf,"conf.ini")
	if err != nil {
		fmt.Println("Failed to parse config file")
	}
	fmt.Println(conf.Section.Enabled)
	fmt.Println(conf.Section.Path)
}

// go get github.com/Unknwon/goconfig  直接读取
func GoConfigReadConf(){
	// 直接获取
	cfg, _ := goconfig.LoadConfigFile("conf.ini")
	enabled,_ := cfg.GetValue("Section","enabled")

	fmt.Println(enabled)
}