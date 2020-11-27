/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/24 下午1:19

*******************************************/
package toml

import (
	"fmt"
	"testing"
)

func TestReadConf(t *testing.T) {
	p,_ := ReadConf("conf.toml")
	fmt.Println(p)
	fmt.Println("==========")
	fmt.Println(p.Db)
}
