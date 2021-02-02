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
<<<<<<< HEAD
	path, _ := cfg.GetValue("Section","path") // 结构体区分大小写，Section，

	fmt.Println(enabled)
	fmt.Println(path)
=======
	path:= cfg.GetKeyComments("Section","path")
	path1,_ := cfg.GetValue("Section","path")

	fmt.Println(enabled)
	fmt.Println(path)
	fmt.Println("path1 = ", path1)
>>>>>>> ac74f244ef88a102760d4ef508919e70688e87de
}
