/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/17 17:21
@Description: 配置文件封装

*********************************************/
package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// 配置文件结构体
type YamlConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Type    string `yaml:"type"`
	Name    string `yaml:"name"`
	Pwrd    string `yaml:"pwrd"`
	Dbname  string `yaml:"dbname"`
	Charset string `yaml:"charset"`
}

func (y *YamlConfig) GetConfig() *YamlConfig {
	yamlFile, err := ioutil.ReadFile("../conf/app.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, y)
	if err != nil {
		fmt.Println(err.Error())
	}
	return y
}
