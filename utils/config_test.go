/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/17 17:56
@Description:

*********************************************/
package utils

import (
	"fmt"
	"testing"
)

func TestYamlConfig_GetConfig(t *testing.T) {
	model := new(YamlConfig)
	confg := model.GetConfig()
	fmt.Println(confg.Name)
}
