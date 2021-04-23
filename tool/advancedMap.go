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
	maps := make([]map[string]interface{}, 3)
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

	fmt.Println("maps=====", maps)
	var outStruct []CardBase

	_ = mapstructure.WeakDecode(maps, &outStruct)

	for k, v := range outStruct {
		fmt.Println("k = ", k, v.Name)
		fmt.Println("k = ", k, v.CardId)
		fmt.Println("k = ", k, v.SalesNum)
		fmt.Println("k = ", k, v.Clicks)
		fmt.Println("========key value =======")
		fmt.Println(outStruct[k].Name)
	}

}


func SileIn (){
	cardId := make([]int,0)
	cardId = append(cardId, 19)
	cardId = append(cardId, 12)
	cardId = append(cardId, 18)
	cardId = append(cardId, 98)

	for _, card := range cardId {
		fmt.Println(card)
	}



}
