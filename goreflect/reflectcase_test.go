/**************************************
 * @Author: Yinjinlin
 * @Description:
 * @File:  reflectcase_test
 * @Version: 1.0.0
 * @Date: 2020/12/18 23:38
 **************************************/
package goreflect

import "testing"

//反射实践案例
func TestStructCase(t *testing.T) {
	var a Monster = Monster{
		Name :"jack ma",
		Age: 18,
		Score: 60.8,
	}
	StructCase(a)
}
