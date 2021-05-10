/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/20 10:19
@Description:

*********************************************/
package tool

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type CardBase struct {
	Name     string  // 名称
	Price    float64 // 价格
	CardId   int     `mapstructure:"card_id"` // 如需要加下划线的，必须自己定义tag
	Clicks   int     // 点击量
	SalesNum int     // 销量 mapstructure 默认映射是小写的
}

type Card struct {
	CardBase
	Sales int // 销量
}

// 定义一个map
func AdvanceMap() {
	maps := make([]map[string]interface{}, 4)

	maps[0] = make(map[string]interface{})
	maps[0]["name"] = "综合0"
	maps[0]["card_id"] = 11
	maps[0]["price"] = 100.0
	maps[0]["salesnum"] = 1003

	maps[1] = make(map[string]interface{})
	maps[1]["name"] = "综合1"
	maps[1]["card_id"] = 12
	maps[1]["price"] = 100.1
	maps[1]["salesnum"] = 1004

	maps[2] = make(map[string]interface{})
	maps[2]["name"] = "综合2"
	maps[2]["card_id"] = 13
	maps[2]["price"] = 200.2
	maps[2]["salesnum"] = 1005

	maps[3] = make(map[string]interface{})
	maps[3]["name"] = "综合3"
	maps[3]["card_id"] = 14
	maps[3]["price"]= 300.5
	maps[3]["salesnum"]=1006

	fmt.Println("maps=====", maps)
	var outStruct []CardBase

	// 追加
	maps = append(maps, map[string]interface{}{
		"name":     "RcardId",
		"card_id":  14,
		"salesnum": 1006,
	})

	fmt.Println("Append maps====",maps)

	fmt.Println("maps[3][\"name\"]",maps[3]["name"])

	maps2 := make([]map[string]interface{}, 2)
	// 必须make，否则panic: assignment to entry in nil map
	// 原因：未初始化的的value 是nil，
	maps2[0] = make(map[string]interface{})
	maps2[0]["name"] = "signle1"
	maps2[0]["card_id"] = 101
	maps2[0]["price"] = 101.0
	maps2[0]["salesnum"] = 1003101

	maps2[1] = make(map[string]interface{})
	maps2[1]["name"] = "signle2"
	maps2[1]["card_id"] = 201
	maps2[1]["price"] = 201.0
	maps2[1]["salesnum"] = 2003101

	maps = append(maps, maps2...)
	fmt.Println("maps =", maps)

	// map 转换成结构体
	fmt.Println("====map 转换成结构体====")
	_ = mapstructure.WeakDecode(maps, &outStruct)

	for k, v := range outStruct {
		fmt.Println("========key value =======")
		fmt.Println("k = ", k, v.Name)
		fmt.Println("k = ", k, v.CardId)
		fmt.Println("k = ", k, v.SalesNum)
		fmt.Println("k = ", k, v.Clicks)

		//fmt.Println(outStruct[k].Name)
	}

}

// 双map
func AdvMapMap(){
	// 第一种初始化
	//maps := map[string]map[string]interface{}{}
	// 第二种初始化
	var maps map[string]map[string]interface{}
	maps = make(map[string]map[string]interface{})
	maps["001"] = make(map[string]interface{})
	maps["001"] = map[string]interface{}{
		"name":"Linux",
	}
	maps["002"] = make(map[string]interface{})
	maps["002"] = map[string]interface{}{
		"cardId":20,
	}
	maps["003"] = make(map[string]interface{})
	maps["003"] = map[string]interface{}{
		"Id": 4,
	}

	fmt.Println(maps)
}

// 切片
func SileIn() {
	cardId := make([]int, 0)
	cardId = append(cardId, 19)
	cardId = append(cardId, 12)
	cardId = append(cardId, 18)
	cardId = append(cardId, 98)
	cardId = append(cardId, 100)

	for _, card := range cardId {
		fmt.Println(card)
	}

	var slice []int
	fmt.Println("slice =", slice)
	slice1 := make([]int, 1)
	fmt.Println("slice1 =", slice1)
	slice2 := make([]int, 0)
	fmt.Println("slice2 =", slice2)



}
