/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/23 14:52
@Description:

*********************************************/
package tool

import (
	"fmt"
	"testing"
)

func TestJsonMarshal(t *testing.T) {
	JsonMarshal()
}

func TestJsonUnmarshal(t *testing.T) {
	JsonUnmarshal()
}

func TestGoStrings(t *testing.T) {

	// 2415 +2655+29490+8505+4680+2712+17205
	// 6465+22365+20745+36000+19755+1100
	fmt.Println(2415 +2655+29490+8505+4680+2712+17205) // 67662
	fmt.Println(6465+22365+20745+36000+19755+1100)  // 106430
}