/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/27 上午10:37

*******************************************/
package base

import "testing"

func TestModel_Where(t *testing.T) {

	var wh Model
	m := map[string]interface{}{"id":1,"name":"jack"}
	wh.Where(m)

	b := []string{"name","age"}
	wh.Field(b)
}

