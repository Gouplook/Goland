/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/24 上午10:42

*******************************************/
package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2" // yaml 库
	"io/ioutil"
)

//profile variables
type YamlConf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
}

func YamlConfTest(){
	var cx YamlConf
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &cx)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(cx)
	fmt.Println(cx.Host)
}
