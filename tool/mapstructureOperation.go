/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/25 下午1:23

*******************************************/
package tool

// go get github.com/mitchellh/mapstructure  库
// map --> struct 的转换
// 使用方法。
// 1: 一般不知道来源的文件是什么个数，先需要进行转换成json格式（encoding/json），map[string]interface{}
// 2: 用json.Unmarshal将字节流解码为map[string]interface{}类型
// 默认情况下，mapstructure 自动映射字段  Name --->name  也可以自己定义： Name ----> usrename

// 注意点 转换时，map与struct 字段名 必须一致。若字段不一样可以用mapstructure:进行更改

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
)

type Person struct {
	Name string `mapstructure:"name"` // 可以映射，默认情况下，mapstructure 自动映射
	Age  int
	// Job  string

}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

// 用法不常用
func MapToStruct() {
	datas := []string{`
    { 
      "type": "person",
      "name":"dj",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

	//	fmt.Println("m=",m)
		switch m["type"].(string) {
		case "person":
			var p Person
			err = mapstructure.Decode(m, &p)
			if err != nil {

			}
			fmt.Println("person", p)

		case "cat":
			var cat Cat
			_ = mapstructure.WeakDecode(m, &cat)
			fmt.Println("cat", cat)
		}
	}

}

// 正向转换
func MapToStruct1() {
	mapdata := map[string]interface{}{
		"name": "jack",
		"age":  19,
		"job":  "goland",
	}
	var structdata Person
	_ = mapstructure.WeakDecode(mapdata, &structdata)
	fmt.Println(structdata)
}

type Person2 struct {
	Name string
	Age  int
	Job  string `mapstructure:"omitempty"`
}

// 反向转换
func StructToMap() {
	p := &Person2{
		Name: "dj",
		Age:  18,
	}
	var m map[string]interface{}
	_ = mapstructure.WeakDecode(p, &m)
	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}
