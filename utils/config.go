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

// mysql配置文件结构体
type YamlConfig struct {
	Host    string `yaml:"host"`    // 主机
	Port    string `yaml:"port"`    // 端口
	Type    string `yaml:"type"`    // 数据库类型
	Name    string `yaml:"name"`    // 数据库主机名
	Pwrd    string `yaml:"pwrd"`    // 数据密码
	Dbname  string `yaml:"dbname"`  // 数据库名字
	Charset string `yaml:"charset"` // 设置字符集
	Maxidle string `yaml:"maxidle"`
	Maxconn string `yaml:"maxconn"` // 最大连接数
}

// 读取配置文件
func (y *YamlConfig) GetConfig() *YamlConfig {
	//yamlFile, err := ioutil.ReadFile("../conf/app.yaml")
	yamlFile, err := ioutil.ReadFile("/Users/uplook/goframe/Goland/conf/app.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	// 序列化位结构体
	err = yaml.Unmarshal(yamlFile, y)
	if err != nil {
		fmt.Println(err.Error())
	}
	return y
}
